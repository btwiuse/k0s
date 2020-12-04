package middleware

import (
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gorilla/handlers"
	"github.com/rs/cors"
)

var (
	GzipMiddleware    = handlers.CompressHandler
	LoggingMiddleware = func(next http.Handler) http.Handler {
		return handlers.LoggingHandler(os.Stderr, next)
	}
	GoroutineMiddleware = func(next http.Handler) http.Handler {
		return http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("You've got", runtime.NumGoroutine(), "goroutines running")
			next.ServeHTTP(w, r)
		}))
	}
    AllowAllCorsMiddleware = func(next http.Handler) http.Handler {
        return cors.AllowAll().Handler(next)
    }
)
