package main

import (
	"log"

	"github.com/silentFellow/neourl/config"
)

func main() {
  // logging set up
  logFile := config.SetupLogging()
  defer logFile.Close()

  log.Println("Just Chilling")
}
