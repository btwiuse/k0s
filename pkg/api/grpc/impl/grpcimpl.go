package impl

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

	"github.com/btwiuse/conntroll/pkg/api"
	"github.com/btwiuse/wetty/pkg/localcmd"
	"github.com/btwiuse/wetty/pkg/msg"
	"github.com/kr/pty"
)

type Session struct {
	*localcmd.Factory
	Name string
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
	lc, err := session.Factory.New()
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
			n, err := lc.Read(buf)
			if err == io.EOF {
				return // nil
			}
			if err != nil {
				return // err
			}
			outputMsg := &api.Message{Type: msg.Type_SESSION_OUTPUT, Body: buf[:n]}
			err = sendServer.Send(outputMsg)
			if err != nil {
				return // err
			}
		}
		return // nil
	}()

	// recv
	for {
		resp, err := sendServer.Recv()
		if err != nil {
			return nil
		}
		log.Println(session.Name, resp.Type, fmt.Sprintf("%q", string(resp.Body)))
		switch resp.Type {
		case msg.Type_CLIENT_INPUT:
			_, err = lc.Write(resp.Body)
			if err != nil {
				log.Println(session.Name, "error writing to lc:", err)
				return err
			}
		case msg.Type_SESSION_RESIZE:
			sz := &pty.Winsize{}
			err = json.Unmarshal(resp.Body, sz)
			if err != nil {
				return err
			}
			err = lc.ResizeTerminal(sz)
			if err != nil {
				return err
			}
		case msg.Type_SESSION_CLOSE:
			return lc.Close()
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
