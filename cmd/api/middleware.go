package api

import (
	"log"
	"net/http"
	"time"
)

func logginMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    start := time.Now()

    next.ServeHTTP(w, r)
    timeTaken := time.Since(start).Seconds()
    log.Printf("Connectted %v-%v took %v seconds.", r.Method, r.URL.Path, timeTaken)
  })
}
