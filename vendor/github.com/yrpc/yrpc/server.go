package yrpc

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"go.uber.org/ratelimit"
	"go.uber.org/zap"
)

var (
	// ErrWriteAfterCloseSelf when try to write after closeself
	ErrWriteAfterCloseSelf = errors.New("write after closeself")
	// ErrRstNonExistingStream when reset non existing stream
	ErrRstNonExistingStream = errors.New("reset non existing stream")
)

// FrameWriter looks like writes a qrpc resp
// but it internally needs be scheduled, thus maintains a simple yet powerful interface
type FrameWriter interface {
	StartWrite(requestID uint64, cmd Cmd, flags FrameFlag)
	WriteBytes(v []byte) // v is copied in WriteBytes
	EndWrite() error     // block until scheduled

	ResetFrame(requestID uint64, reason Cmd) error
}

// StreamWriter is returned by StreamRequest
type StreamWriter interface {
	RequestID() uint64
	StartWrite(cmd Cmd)
	WriteBytes(v []byte)     // v is copied in WriteBytes
	EndWrite(end bool) error // block until scheduled
}

// A Handler responds to an qrpc request.
type Handler interface {
	// FrameWriter will be recycled when ServeQRPC finishes, so don't cache it
	ServeQRPC(FrameWriter, *RequestFrame)
}

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as qrpc handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.
type HandlerFunc func(FrameWriter, *RequestFrame)

// ServeQRPC calls f(w, r).
func (f HandlerFunc) ServeQRPC(w FrameWriter, r *RequestFrame) {
	f(w, r)
}

// ServeMux is qrpc request multiplexer.
type ServeMux struct {
	mu sync.RWMutex
	m  map[Cmd]Handler
}

// NewServeMux allocates and returns a new ServeMux.
func NewServeMux() *ServeMux { return new(ServeMux) }

// HandleFunc registers the handler function for the given pattern.
func (mux *ServeMux) HandleFunc(cmd Cmd, handler func(FrameWriter, *RequestFrame)) {
	mux.Handle(cmd, HandlerFunc(handler))
}

// Handle registers the handler for the given pattern.
// If a handler already exists for pattern, handle panics.
func (mux *ServeMux) Handle(cmd Cmd, handler Handler) {
	mux.mu.Lock()
	defer mux.mu.Unlock()

	if handler == nil {
		panic("qrpc: nil handler")
	}
	if _, exist := mux.m[cmd]; exist {
		panic("qrpc: multiple registrations for " + string(cmd))
	}

	if mux.m == nil {
		mux.m = make(map[Cmd]Handler)
	}
	mux.m[cmd] = handler
}

// ServeQRPC dispatches the request to the handler whose
// cmd matches the request.
func (mux *ServeMux) ServeQRPC(w FrameWriter, r *RequestFrame) {
	mux.mu.RLock()
	h, ok := mux.m[r.Cmd]
	if !ok {
		l.Error("cmd not registered", zap.Uint32("cmd", uint32(r.Cmd)))
		r.Close()
		return
	}
	mux.mu.RUnlock()
	h.ServeQRPC(w, r)
}

// Server defines parameters for running an qrpc server.
type Server struct {
	// one handler for each listening address
	conf   ServerConfig
	upTime time.Time

	// manages below
	mu           sync.Mutex
	listener     net.Listener
	doneChan     chan struct{}
	shutdownFunc func()
	done         bool

	id2Conn          sync.Map
	activeConn       sync.Map // for better iterate when write, map[*serveconn]struct{}
	closeRateLimiter ratelimit.Limiter

	wg sync.WaitGroup // wait group for goroutines

	pushID uint64
}

// NewServer creates a server
func NewServer(conf ServerConfig) *Server {
	var closeRateLimiter ratelimit.Limiter
	if conf.MaxCloseRate != 0 {
		closeRateLimiter = ratelimit.New(conf.MaxCloseRate)
	}
	if conf.WriteFrameChSize < 1 {
		// at least 1 for WriteFrameChSize
		conf.WriteFrameChSize = 1
	}
	return &Server{
		conf:             conf,
		upTime:           time.Now(),
		doneChan:         make(chan struct{}),
		closeRateLimiter: closeRateLimiter,
	}
}

// ListenAndServe starts listening on all conf
func (srv *Server) ListenAndServe() (err error) {
	var (
		rawln net.Listener
		yln   net.Listener
	)

	if srv.conf.ListenFunc != nil {
		rawln, err = srv.conf.ListenFunc("tcp", srv.conf.Addr)
	} else {
		rawln, err = net.Listen("tcp", srv.conf.Addr)
	}
	if err != nil {
		return err
	}

	if srv.conf.OverlayNetwork != nil {
		yln = srv.conf.OverlayNetwork(rawln)
	} else {
		yln = rawln
	}

	return srv.Serve(yln)
}

// BindingConfig for retrieve ServerConfig
func (srv *Server) BindingConfig() ServerConfig {
	return srv.conf
}

var (

	// ErrServerClosed is returned by the Server's Serve, ListenAndServe,
	// methods after a call to Shutdown or Close.
	ErrServerClosed = errors.New("qrpc: Server closed")
	// ErrListenerAcceptReturnType when Listener.Accept doesn't return TCPConn
	ErrListenerAcceptReturnType = errors.New("qrpc: Listener.Accept doesn't return TCPConn")
	// ErrAcceptTimedout when accept timed out
	ErrAcceptTimedout = errors.New("qrpc: accept timed out")

	defaultAcceptTimeout = 5 * time.Second
)

// Serve accepts incoming connections on the Listener qrpcListener, creating a
// new service goroutine for each. The service goroutines read requests and
// then call srv.Handler to reply to them.
//
// Serve always returns a non-nil error. After Shutdown or Close, the
// returned error is ErrServerClosed.
func (srv *Server) Serve(ln net.Listener) error {
	var tempDelay time.Duration // how long to sleep on accept failure

	serveCtx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	for {
		rw, e := ln.Accept()
		if e != nil {
			select {
			case <-srv.doneChan:
				return ErrServerClosed
			default:
			}
			if e == ErrAcceptTimedout {
				// for overlay network
				continue
			}
			if opError, ok := e.(*net.OpError); ok && opError.Timeout() {
				// don't log the scheduled timeout
				continue
			}
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				l.Error("qrpc: Accept", zap.Duration("retrying in", tempDelay), zap.Error(e))
				time.Sleep(tempDelay)
				continue
			}
			l.Error("qrpc: Accept fatal", zap.Error(e)) // accept4: too many open files in system
			time.Sleep(time.Second)                     // keep trying instead of quit
			continue
		}
		tempDelay = 0

		GoFunc(&srv.wg, func() {
			c := srv.newConn(serveCtx, rw)
			c.serve()
		})
	}
}

// Create new connection from rwc.
func (srv *Server) newConn(ctx context.Context, rwc net.Conn) (sc *serveconn) {
	sc = &serveconn{
		server:         srv,
		rwc:            rwc,
		untrackedCh:    make(chan struct{}),
		cs:             &ConnStreams{},
		readFrameCh:    make(chan readFrameResult, srv.conf.ReadFrameChSize),
		writeFrameCh:   make(chan *writeFrameRequest, srv.conf.WriteFrameChSize),
		cachedRequests: make([]*writeFrameRequest, 0, srv.conf.WriteFrameChSize),
		cachedBuffs:    make(net.Buffers, 0, srv.conf.WriteFrameChSize),
		wlockCh:        make(chan struct{}, 1)}

	ctx, cancelCtx := context.WithCancel(ctx)
	ctx = context.WithValue(ctx, ConnectionInfoKey, &ConnectionInfo{SC: sc})

	sc.cancelCtx = cancelCtx
	sc.ctx = ctx
	sc.bytesWriter = NewWriterWithTimeout(ctx, rwc, srv.conf.DefaultWriteTimeout)

	srv.activeConn.Store(sc, struct{}{})

	return sc
}

var kickOrder uint64

// bindID bind the id to sc
// it is concurrent safe
func (srv *Server) bindID(sc *serveconn, id string) (kick bool, ko uint64) {

check:
	v, loaded := srv.id2Conn.LoadOrStore(id, sc)

	if loaded {
		vsc := v.(*serveconn)
		if vsc == sc {
			return
		}
		ok, ch := srv.untrack(vsc, true)
		if !ok {
			<-ch
		}
		l.Debug("trigger closeUntracked", zap.Uintptr("sc", uintptr(unsafe.Pointer(sc))), zap.Uintptr("vsc", uintptr(unsafe.Pointer(vsc))))

		err := vsc.closeUntracked()
		if err != nil {
			if opErr, ok := err.(*net.OpError); ok {
				err = opErr.Err
			}
		}

		if srv.conf.CounterMetric != nil {
			errStr := fmt.Sprintf("%v", err)
			countlvs := []string{"method", "kickoff", "error", errStr}
			srv.conf.CounterMetric.With(countlvs...).Add(1)
		}

		atomic.AddUint64(&kickOrder, 1)
		kick = true

		goto check
	}

	ko = atomic.LoadUint64(&kickOrder)
	return
}

func (srv *Server) untrack(sc *serveconn, kicked bool) (bool, <-chan struct{}) {

	locked := atomic.CompareAndSwapUint32(&sc.untrack, 0, 1)
	if !locked {
		return false, sc.untrackedCh
	}

	id := sc.GetID()
	if id != "" {
		srv.id2Conn.Delete(id)
	}
	srv.activeConn.Delete(sc)

	if kicked {
		if srv.conf.OnKickCB != nil {
			srv.conf.OnKickCB(sc.GetWriter())
		}
	}
	close(sc.untrackedCh)
	return true, sc.untrackedCh
}

var shutdownPollInterval = 500 * time.Millisecond

// Shutdown gracefully shutdown the server
func (srv *Server) Shutdown() error {

	srv.mu.Lock()
	if srv.done {
		srv.mu.Unlock()
		goto done
	}

	{
		lnerr := srv.closeListenersLocked()
		if lnerr != nil {
			srv.mu.Unlock()
			return lnerr
		}
	}

	srv.done = true
	srv.mu.Unlock()

	close(srv.doneChan)

	if srv.shutdownFunc != nil {
		srv.shutdownFunc()
	}

done:
	srv.wg.Wait()

	return nil
}

// OnShutdown registers f to be called when shutdown
func (srv *Server) OnShutdown(f func()) {

	srv.mu.Lock()
	if srv.done {
		srv.mu.Unlock()
		f()
		return
	}

	srv.shutdownFunc = f
	srv.mu.Unlock()
}

// GetPushID gets the pushId
func (srv *Server) GetPushID() uint64 {
	pushID := atomic.AddUint64(&srv.pushID, 1)
	return pushID
}

// WalkConnByID iterates over  serveconn by ids
func (srv *Server) WalkConnByID(ids []string, f func(FrameWriter, *ConnectionInfo)) {
	for _, id := range ids {
		v, ok := srv.id2Conn.Load(id)
		if ok {
			sc := v.(*serveconn)
			f(v.(*serveconn).GetWriter(), sc.ctx.Value(ConnectionInfoKey).(*ConnectionInfo))
		}
	}
}

// GetConnectionInfoByID returns the ConnectionInfo for id
func (srv *Server) GetConnectionInfoByID(id string) *ConnectionInfo {
	v, ok := srv.id2Conn.Load(id)
	if !ok {
		return nil
	}

	return v.(*serveconn).ctx.Value(ConnectionInfoKey).(*ConnectionInfo)
}

// WalkConn walks through each serveconn
func (srv *Server) WalkConn(f func(FrameWriter, *ConnectionInfo) bool) {
	srv.activeConn.Range(func(k, v interface{}) bool {
		sc := k.(*serveconn)
		return f(sc.GetWriter(), sc.ctx.Value(ConnectionInfoKey).(*ConnectionInfo))
	})
}

func (srv *Server) closeListenersLocked() (err error) {
	ln := srv.listener
	if ln == nil {
		return nil
	}
	if err = ln.Close(); err != nil {
		return
	}
	srv.listener = nil
	return nil
}
