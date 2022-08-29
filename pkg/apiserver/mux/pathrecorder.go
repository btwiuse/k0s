package mux

import (
	"net/http"
	"sort"
	"strings"
)

// https://github.dev/kubernetes/apiserver/tree/master/pkg/server

type PathRecorderMux struct {
	name            string
	notFoundHandler http.Handler
	pathToHandler   map[string]http.Handler
	prefixToHandler map[string]http.Handler
	exposedPaths    []string
}

func (m *PathRecorderMux) NotFoundHandler(notFoundHandler http.Handler) {
	m.notFoundHandler = notFoundHandler
}

func (m *PathRecorderMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if exactHandler, ok := m.pathToHandler[r.URL.Path]; ok {
		exactHandler.ServeHTTP(w, r)
		return
	}

	for prefix, prefixHandler := range m.prefixToHandler {
		if strings.HasPrefix(r.URL.Path, prefix) {
			prefixHandler.ServeHTTP(w, r)
			return
		}
	}

	m.notFoundHandler.ServeHTTP(w, r)
}

func (m *PathRecorderMux) UnlistedHandle(path string, handler http.Handler) {
	m.pathToHandler[path] = handler
}

func NewPathRecorderMux(name string) *PathRecorderMux {
	ret := &PathRecorderMux{
		name:            name,
		pathToHandler:   map[string]http.Handler{},
		prefixToHandler: map[string]http.Handler{},
		exposedPaths:    []string{},
	}

	return ret
}

// ListedPaths returns the registered handler exposedPaths.
func (m *PathRecorderMux) ListedPaths() []string {
	handledPaths := append([]string{}, m.exposedPaths...)
	sort.Strings(handledPaths)

	return handledPaths
}
