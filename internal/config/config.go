package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	ENV  string
}

func MustLoad() Config {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == ""{
		panic("PORT is req");
	}

	env := os.Getenv("ENV")
	if env == ""{
		panic("ENV is not set, defaulting to development")
	}

	return Config{
		Port: port,
		ENV:  env,
	}
}