package initialization

import (
	"os"

	"github.com/go-chi/cors"
)

func LoadConfig() (*cors.Options, string, string, string, error) {
	serverPort := os.Getenv("SERVER_PORT")
	logLevel := os.Getenv("LOG_LEVEL")
	logFormat := os.Getenv("LOG_FORMAT")

	const maxAge = 300
	return &cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "x-api-key"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           maxAge,
	}, serverPort, logLevel, logFormat, nil
}
