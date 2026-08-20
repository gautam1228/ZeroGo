package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Env  string
}

// This function panics, instead of returning an error, since we start it with the prefix of Must
func MustLoad() Config {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading the .env file.")
	}

	env := os.Getenv("ENV")
	if env == "" {
		log.Fatalf("ENV is required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("PORT is required")
	}

	return Config{
		Port: port,
		Env:  env,
	}
}
