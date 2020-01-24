package agent

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

var (
	GzipMiddleware    = handlers.CompressHandler
	LoggingMiddleware = func(next http.Handler) http.Handler {
		return handlers.LoggingHandler(os.Stderr, next)
	}
)
