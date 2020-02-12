package yrpc

import (
	"bufio"
	"context"
	"encoding/binary"
	"io"
	"net"
	"sync"
)

// Reader read data from socket
type Reader struct {
	conn   net.Conn
	reader *bufio.Reader
	ctx    context.Context
}

const (
	// ReadBufSize for read buf
	ReadBufSize = 1024
)

var bufPool = sync.Pool{New: func() interface{} {
	return bufio.NewReaderSize(nil, ReadBufSize)
}}

// NewReader creates a StreamReader instance
func NewReader(ctx context.Context, conn net.Conn) *Reader {
	if ctx == nil {
		ctx = context.Background()
	}

	bufReader := bufPool.Get().(*bufio.Reader)
	bufReader.Reset(conn)
	return &Reader{ctx: ctx, conn: conn, reader: bufReader}
}

// Finalize is called when no longer used
func (r *Reader) Finalize() {
	r.reader.Reset(nil)
	bufPool.Put(r.reader)
	r.reader = nil
}

// ReadUint32 read uint32 from socket
func (r *Reader) ReadUint32() (uint32, error) {
	bytes := make([]byte, 4)
	err := r.ReadBytes(bytes)
	if err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint32(bytes), nil
}

// ReadBytes read bytes honouring CtxCheckMaxInterval
func (r *Reader) ReadBytes(bytes []byte) (err error) {
	var (
		offset int
		n      int
	)
	size := len(bytes)

	for {
		n, err = io.ReadFull(r.reader, bytes[offset:])
		offset += n
		if err != nil {
			return err
		}
		if offset >= size {
			return nil
		}

		select {
		case <-r.ctx.Done():
			return r.ctx.Err()
		default:
		}
	}

}
