package api

import (
	"log"
	"net/http"
	"time"
)

type middleware func(http.Handler) http.Handler

func createMiddlewareStack(xs ...middleware) middleware {
  return func(next http.Handler) http.Handler {
    for i:= len(xs) - 1; i >= 0; i-- {
      x := xs[i]
      next = x(next)
    }

    return next
  }
}

func loggingMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    start := time.Now()

    next.ServeHTTP(w, r)
    timeTaken := time.Since(start).Seconds()
    log.Printf("Connectted %v-%v took %v seconds.", r.Method, r.URL.Path, timeTaken)
  })
}

func corsMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Replace "*" with specific origins if needed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "false") // Set to "true" if credentials are required

    if r.Method == http.MethodOptions {
      w.WriteHeader(http.StatusNoContent)
      return
    }

    next.ServeHTTP(w, r)
  })
}
