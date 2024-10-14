package config

import (
	"os"
)

type Config struct {
	PublicHost string
	Port       string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		PublicHost: os.Getenv("PUBLIC_HOST"),
		Port:       os.Getenv("PORT"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
