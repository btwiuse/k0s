package apiserver

import (
	"net/http"
	"strings"

	"github.com/emicklei/go-restful/v3"
	"k0s.io/pkg/apiserver/mux"
)

type APIServerHandler struct {
	FullHandlerChain http.Handler

	GoRestfulContainer *restful.Container

	NonGoRestfulMux *mux.PathRecorderMux

	Director http.Handler
}

func (a *APIServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.FullHandlerChain.ServeHTTP(w, r)
}

type HandlerChainBuilderFn func(apiHandler http.Handler) http.Handler

func DefaultAPIServerHandler() *APIServerHandler {
	var (
		name                string                = "default"
		handlerChainBuilder HandlerChainBuilderFn = func(x http.Handler) http.Handler { return x }
		notFoundHandler     http.Handler          = http.NotFoundHandler()
	)

	return NewAPIServerHandler(name, handlerChainBuilder, notFoundHandler)
}

func NewAPIServerHandler(name string, handlerChainBuilder HandlerChainBuilderFn, notFoundHandler http.Handler) *APIServerHandler {
	nonGoRestfulMux := mux.NewPathRecorderMux(name)

	if notFoundHandler != nil {
		nonGoRestfulMux.NotFoundHandler(notFoundHandler)
	}

	gorestfulContainer := restful.NewContainer()
	gorestfulContainer.ServeMux = http.NewServeMux()
	gorestfulContainer.Router(restful.CurlyRouter{}) // e.g. for proxy/{kind}/{name}/{*}
	// gorestfulContainer.RecoverHandler
	// gorestfulContainer.ServiceErrorHandler

	director := director{
		name:               name,
		goRestfulContainer: gorestfulContainer,
		nonGoRestfulMux:    nonGoRestfulMux,
	}

	return &APIServerHandler{
		FullHandlerChain:   handlerChainBuilder(director),
		GoRestfulContainer: gorestfulContainer,
		NonGoRestfulMux:    nonGoRestfulMux,
		Director:           director,
	}
}

type director struct {
	name               string
	goRestfulContainer *restful.Container
	nonGoRestfulMux    *mux.PathRecorderMux
}

func goRestfulMatches(path string, ws *restful.WebService) bool {
	switch {
	case (ws.RootPath() == "/apis"):
		return path == "/apis" || path == "/apis/"
	case strings.HasPrefix(path, ws.RootPath()):
		// ensure an exact match or a path boundary match
		return len(path) == len(ws.RootPath()) || path[len(ws.RootPath())] == '/'
	default:
		return false
	}
}

func (d director) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, ws := range d.goRestfulContainer.RegisteredWebServices() {
		if goRestfulMatches(r.URL.Path, ws) {
			d.goRestfulContainer.Dispatch(w, r)
			return
		}
	}

	d.nonGoRestfulMux.ServeHTTP(w, r)
}
