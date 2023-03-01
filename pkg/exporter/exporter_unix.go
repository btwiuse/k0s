//go:build !windows && !android && !linux

package exporter

import (
	"io"
	"net/http"
)

func NewHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "# /metrics is not implemented for unix platform")
	})
}
