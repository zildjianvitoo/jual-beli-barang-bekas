package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
}

func SetupEnv() (cfg AppConfig, err error) {

	godotenv.Load(".env")

	httpPort := os.Getenv("HTTP_PORT")
	if len(httpPort) == 0 {
		return AppConfig{}, errors.New("env variable not found")
	}

	return AppConfig{ServerPort: httpPort}, nil
}
