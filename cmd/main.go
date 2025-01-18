package main

import (
	"log"
	"net/http"

	"github.com/silentFellow/neourl/cmd/api"
	"github.com/silentFellow/neourl/config"
	"github.com/silentFellow/neourl/internal/urlcoder"
)

func main() {
  // logging set up
  logFile := config.SetupLogging()
  defer logFile.Close()

  // map storage set up
  urlStorage := urlcoder.NewURLStorage()

  // setting server up
  PORT := config.Envs.Server_Port
  server := api.NewServer(PORT, urlStorage)

  // GracefulShutdown
  done := make(chan struct{})
  go api.GracefulShutdown(server, done)

  // running server
  log.Println("Server started at port: ", PORT)
  err := server.Run()
  if err != nil && err != http.ErrServerClosed {
    log.Fatal("Failed to start the server", err)
  }

  <- done
  log.Println("Application terminated successfully")
}
