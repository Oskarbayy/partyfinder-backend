package router

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter()

	// Global middleware
	r.Use(loggingMiddleware)
	r.Use(recoveryMiddleware)
	r.Use(timeoutMiddleware(60 * time.Second))

	return r
}

// loggingMiddleware logs each request’s method, path, and duration.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s took %v", r.Method, r.URL.Path, time.Since(start))
	})
}

// recoveryMiddleware recovers from panics and returns a 500.
func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %v", err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// timeoutMiddleware aborts any request taking longer than d.
func timeoutMiddleware(d time.Duration) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.TimeoutHandler(next, d, `{"error":"timeout"}`)
	}
}
