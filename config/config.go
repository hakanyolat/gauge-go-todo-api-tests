package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	TestUrl        string
	SeleniumHubUrl string
}

func GetConfig() Config {
	if os.Getenv("GOPATH") == "" {
		panic("GOPATH not found in environment")
	}

	if err := godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/hakanyolat/gauge-go-todo-api-tests/.env")); err != nil {
		panic("Error loading .env file.")
	}

	return Config{
		TestUrl:        os.Getenv("GAUGE_TEST_URL"),
		SeleniumHubUrl: os.Getenv("GAUGE_SELENIUM_HUB_URL"),
	}
}
