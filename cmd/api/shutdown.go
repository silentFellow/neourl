package api

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"
)

func GracefulShutdown(s *server, done chan struct{}) {
  ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
  defer cancel()

  <- ctx.Done()

  log.Println("Server going to shutdown within 5 seconds, CTRL+C to force quit")
  ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()

  if err := s.server.Shutdown(ctx); err != nil {
    log.Fatal("Failed to close server gracefully")
  }
  
  log.Println("Exiting server")
  done <- struct{}{}
}
