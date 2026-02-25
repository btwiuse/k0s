package asciitransport

import (
	"io"
	"net"
	"testing"
	"time"
)

func TestServerSessionReadWrite(t *testing.T) {
	clientConn, serverConn := net.Pipe()
	defer clientConn.Close()

	session := Server(serverConn)
	defer session.Close()

	// Client sends an input event: [0, "i", "hello"]
	inputMsg := `[0,"i","hello"]` + "\n"
	go func() {
		clientConn.Write([]byte(inputMsg))
	}()

	// Server reads the input via io.Reader
	buf := make([]byte, 1024)
	n, err := session.Read(buf)
	if err != nil {
		t.Fatalf("Read error: %v", err)
	}
	if got := string(buf[:n]); got != "hello" {
		t.Fatalf("Read: got %q, want %q", got, "hello")
	}
}

func TestServerSessionWrite(t *testing.T) {
	clientConn, serverConn := net.Pipe()
	defer clientConn.Close()

	session := Server(serverConn)
	defer session.Close()

	// Server writes output via io.Writer
	go func() {
		session.Write([]byte("world"))
	}()

	// Client reads the output event from the connection
	buf := make([]byte, 4096)
	n, err := clientConn.Read(buf)
	if err != nil {
		t.Fatalf("clientConn.Read error: %v", err)
	}
	line := string(buf[:n])
	// The output should be JSON formatted like [0,"o","world"]
	if len(line) == 0 {
		t.Fatal("expected non-empty output")
	}
	// Verify it contains the output data
	if !containsSubstring(line, "world") {
		t.Fatalf("output %q does not contain %q", line, "world")
	}
}

func TestServerSessionNextResize(t *testing.T) {
	clientConn, serverConn := net.Pipe()
	defer clientConn.Close()

	session := Server(serverConn)
	defer session.Close()

	// Client sends a resize event (JSON object)
	resizeMsg := `{"version":2,"width":120,"height":40}` + "\n"
	go func() {
		clientConn.Write([]byte(resizeMsg))
	}()

	re, err := session.NextResize()
	if err != nil {
		t.Fatalf("NextResize error: %v", err)
	}
	if re.Width != 120 {
		t.Fatalf("Width: got %d, want 120", re.Width)
	}
	if re.Height != 40 {
		t.Fatalf("Height: got %d, want 40", re.Height)
	}
}

func TestServerSessionClose(t *testing.T) {
	clientConn, serverConn := net.Pipe()
	defer clientConn.Close()

	session := Server(serverConn)

	// Close should return nil
	if err := session.Close(); err != nil {
		t.Fatalf("Close error: %v", err)
	}

	// After close, Read should return EOF
	buf := make([]byte, 1024)
	done := make(chan struct{})
	go func() {
		_, err := session.Read(buf)
		if err != io.EOF {
			t.Errorf("Read after Close: got %v, want io.EOF", err)
		}
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("Read after Close did not return within timeout")
	}
}

func TestServerSessionNextResizeAfterClose(t *testing.T) {
	clientConn, serverConn := net.Pipe()
	defer clientConn.Close()

	session := Server(serverConn)
	session.Close()

	done := make(chan struct{})
	go func() {
		_, err := session.NextResize()
		if err != io.EOF {
			t.Errorf("NextResize after Close: got %v, want io.EOF", err)
		}
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("NextResize after Close did not return within timeout")
	}
}

func containsSubstring(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && stringContains(s, substr))
}

func stringContains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
