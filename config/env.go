package config

import (
	"os"

	"github.com/joho/godotenv"
)

type env struct {
  Server_Port string
}

var Envs env = initEnv()

func initEnv() env {
  return env {
    Server_Port: getEnv("SERVER_PORT", "8080"),
  }
}

func getEnv(key, fallback string) string {
  godotenv.Load()

  if val, ok := os.LookupEnv(key); ok {
    return val
  }

  return fallback
}
