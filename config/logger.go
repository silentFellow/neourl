package config

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
  LOG_DIR = "logs"
  LOG_FILE = "server.log"
)

func SetupLogging() *os.File {
  err := os.MkdirAll(LOG_DIR, 0777)
  if err != nil {
    log.Fatalln("Failed to create log dir", err)
  }

  logPath := fmt.Sprintf("%v/%v", LOG_DIR, LOG_FILE)
  logFile, err := os.OpenFile(logPath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0777)
  if err != nil {
    log.Fatalln("Failed to create log file", err)
  }

  writter := io.MultiWriter(logFile, os.Stdout)
  log.SetOutput(writter)

  return logFile
}
