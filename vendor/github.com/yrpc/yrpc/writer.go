package yrpc

import (
	"context"
	"net"
	"time"
)

// Writer writes data to connection
type Writer struct {
	ctx  context.Context
	conn net.Conn
}

const (
	// CtxCheckMaxInterval for check ctx.Done
	CtxCheckMaxInterval = 3 * time.Second
)

// NewWriter new instance
func NewWriter(ctx context.Context, conn net.Conn) *Writer {
	return &Writer{ctx: ctx, conn: conn}
}

// Write writes bytes
func (w *Writer) Write(bytes []byte) (int, error) {
	var (
		offset int
		n      int
		err    error
	)

	size := len(bytes)

	for {
		n, err = w.conn.Write(bytes[offset:])
		offset += n
		if err != nil {
			return offset, err
		}
		if offset >= size {
			return offset, nil
		}

		select {
		case <-w.ctx.Done():
			return offset, w.ctx.Err()
		default:
		}
	}

}

func (w *Writer) writeBuffers(buffs *net.Buffers) (int64, error) {
	var (
		offset int64
		n      int64
		err    error
	)

	var size int64
	for _, bytes := range *buffs {
		size += int64(len(bytes))
	}

	for {
		n, err = buffs.WriteTo(w.conn)
		offset += n
		if err != nil {
			return offset, err
		}
		if offset >= size {
			return offset, nil
		}

		select {
		case <-w.ctx.Done():
			return offset, w.ctx.Err()
		default:
		}
	}
}
