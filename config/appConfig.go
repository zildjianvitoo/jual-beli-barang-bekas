package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort     string
	DatasourceName string
	AppSecret      string
}

func SetupEnv() (cfg AppConfig, err error) {

	godotenv.Load(".env")

	httpPort := os.Getenv("HTTP_PORT")
	if len(httpPort) == 0 {
		return AppConfig{}, errors.New("env variable not found")
	}

	datasourceName := os.Getenv("DATASOURCE_NAME")
	if len(datasourceName) == 0 {
		return AppConfig{}, errors.New("env variable not found")
	}

	appSecret := os.Getenv("APP_SECRET")
	if len(appSecret) == 0 {
		return AppConfig{}, errors.New("env variable not found")
	}

	return AppConfig{ServerPort: httpPort, DatasourceName: datasourceName, AppSecret: appSecret}, nil
}
