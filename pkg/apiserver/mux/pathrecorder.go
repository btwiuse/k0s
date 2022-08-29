package mux

import (
	"net/http"
)

// https://github.dev/kubernetes/apiserver/tree/master/pkg/server

type PathRecorderMux struct {
	name            string
	notFoundHandler http.Handler
}

func (m *PathRecorderMux) NotFoundHandler(notFoundHandler http.Handler) {
	m.notFoundHandler = notFoundHandler
}

func (m *PathRecorderMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.notFoundHandler.ServeHTTP(w, r)
}

func NewPathRecorderMux(name string) *PathRecorderMux {
	return &PathRecorderMux{}
}
