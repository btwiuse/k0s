package grpcimpl

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strings"

	"github.com/btwiuse/conntroll/pkg/agent"
	"github.com/btwiuse/conntroll/pkg/api"
	"github.com/btwiuse/wetty/pkg/msg"
)

type Session struct {
	TtyFactory agent.TtyFactory
	// client id/index, to distinguish logs of different commands
}

func (session *Session) Chunker(req *api.ChunkRequest, chunkerServer api.Session_ChunkerServer) error {
	// req.Path
	// req.Request
	// log.Println("Chunker called with", req)

	// no such file => fileserver
	// isdir => fileserver
	// openerror => fileserver
	// other => chunker

	var (
		path = filepath.Clean(req.Path)

		statfail bool
		openfail bool
		isdir    bool
		issmall  bool

		filename string
		filesize int64

		reader io.Reader
	)

	info, staterr := os.Stat(path)
	if staterr == nil {
		if !info.IsDir() {
			filesize = info.Size()
			if filesize > 4*(1<<20) { // 4M
				f, openerr := os.Open(req.Path)
				if openerr == nil {
					filename = filepath.Base(f.Name())
					defer f.Close()
					header := fmt.Sprintf(ResponseHeaderTemplate, filename, filesize)
					h := strings.NewReader(header)
					reader = io.MultiReader(h, f)
				} else {
					openfail = true
				}
			} else {
				issmall = true
			}
		} else {
			isdir = true
		}
	} else {
		statfail = true
	}

	if statfail || openfail || isdir || issmall {
		w := httptest.NewRecorder()
		r, err := http.ReadRequest(bufio.NewReader(bytes.NewBuffer(req.Request)))
		if err != nil {
			return err
		}
		http.FileServer(http.Dir("/")).ServeHTTP(w, r)
		resp, err := httputil.DumpResponse(w.Result(), true)
		if err != nil {
			return err
		}
		// fmt.Printf("%s", resp)
		reader = bytes.NewReader(resp)
	} else {
		log.Println("Chunker!!!", filename)
	}

	buf := make([]byte, 64*1024)
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		chunk := &api.Chunk{
			Chunk: buf[:n],
		}
		err = chunkerServer.Send(chunk)
		if err != nil {
			return err
		}
	}
	return nil
}

func (session *Session) Send(sendServer api.Session_SendServer) error {
	tty, err := session.TtyFactory.MakeTty()
	if err != nil {
		return err
	}

	// send
	go func() {
		buf := make([]byte, 1<<12-1)
		if err != nil {
			return // err
		}
		for {
			n, err := tty.Read(buf)
			if err == io.EOF {
				return // nil
			}
			if err != nil {
				return // err
			}
			var (
				msgType = msg.Type_SESSION_OUTPUT
				msgBody = buf[:n]
				req = &api.Message{
					Type: msgType,
					Body: msgBody,
				}
			)
			err = sendServer.Send(req)
			if err != nil {
				return // err
			}
			log.Println(req.Type, fmt.Sprintf("%q", string(req.Body)))
		}
		return // nil
	}()

	// recv
	for {
		resp, err := sendServer.Recv()
		if err != nil {
			return nil
		}
		log.Println(resp.Type, fmt.Sprintf("%q", string(resp.Body)))
		switch resp.Type {
		case msg.Type_CLIENT_INPUT:
			_, err = tty.Write(resp.Body)
			if err != nil {
				log.Println("error writing to tty:", err)
				return err
			}
		case msg.Type_SESSION_RESIZE:
			type Winsize struct {
				Rows int
				Cols int
			}
			sz := &Winsize{}
			err = json.Unmarshal(resp.Body, sz)
			if err != nil {
				return err
			}
			err = tty.Resize(sz.Rows, sz.Cols)
			if err != nil {
				return err
			}
		case msg.Type_SESSION_CLOSE:
			return tty.Close()
		}
	}

	return nil
}

var ResponseHeaderTemplate = strings.Join([]string{
	"HTTP/1.1 200 OK",
	"Accept-Ranges: bytes",
	"Content-Type: application/octet-stream",
	"Content-Disposition: attachment; filename=%q",
	"Last-Modified: Sun, 29 Sep 2019 03:58:56 GMT",
	"Content-Length: %d",
	"\r\n"}, "\r\n")
