package server

import (
	"net"
	"testing"
)

func TestChannelListenerAcceptAndClose(t *testing.T) {
	cl := NewChannelListener()

	// Send a connection through the channel
	go func() {
		c1, c2 := net.Pipe()
		defer c2.Close()
		cl.Conns <- c1
	}()

	conn, err := cl.Accept()
	if err != nil {
		t.Fatalf("Accept() returned unexpected error: %v", err)
	}
	conn.Close()

	// Close should not panic
	if err := cl.Close(); err != nil {
		t.Fatalf("Close() returned unexpected error: %v", err)
	}

	// Accept after close should return net.ErrClosed
	conn, err = cl.Accept()
	if err != net.ErrClosed {
		t.Fatalf("Accept() after Close() returned error %v, want net.ErrClosed", err)
	}
	if conn != nil {
		t.Fatalf("Accept() after Close() returned non-nil conn")
	}
}

func TestChannelListenerDoubleClose(t *testing.T) {
	cl := NewChannelListener()

	// First close should succeed
	if err := cl.Close(); err != nil {
		t.Fatalf("first Close() returned unexpected error: %v", err)
	}

	// Second close should not panic
	if err := cl.Close(); err != nil {
		t.Fatalf("second Close() returned unexpected error: %v", err)
	}
}

func TestChannelListenerAddr(t *testing.T) {
	cl := NewChannelListener()
	defer cl.Close()

	if cl.Network() != "hijack" {
		t.Fatalf("Network() = %q, want %q", cl.Network(), "hijack")
	}
	if cl.String() != "hijack" {
		t.Fatalf("String() = %q, want %q", cl.String(), "hijack")
	}
	if cl.Addr() != cl {
		t.Fatal("Addr() should return the listener itself")
	}
}
