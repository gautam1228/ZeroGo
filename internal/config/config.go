package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	Env         string
	DatabaseUrl string
}

// This function panics, instead of returning an error, since we start it with the prefix of Must
func MustLoad() Config {
	godotenv.Load()

	env := os.Getenv("ENV")
	if env == "" {
		panic("ENV is required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT is required")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		panic("DB_URL is required")
	}

	return Config{
		Port:        port,
		Env:         env,
		DatabaseUrl: dbUrl,
	}
}
