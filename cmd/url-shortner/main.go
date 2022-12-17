package main

import (
	"github.com/keivanipchihagh/url-shortener/api"
	"github.com/keivanipchihagh/url-shortener/internal/dotenv"
	"github.com/keivanipchihagh/url-shortener/internal/postgres"
)

func main() {

	dotenv.LoadEnv()

	// Initialize the database client parameters
	postgres.InitializeClient(
		dotenv.ReadEnv("POSTGRES_HOST"),
		dotenv.ReadEnvAsInt("POSTGRES_PORT"),
		dotenv.ReadEnv("POSTGRES_USER"),
		dotenv.ReadEnv("POSTGRES_PASSWORD"),
		dotenv.ReadEnv("POSTGRES_DB"),
	)

	// Initialize the router and start the API server
	api.RegisterRoutes()
	api.Run("0.0.0.0", dotenv.ReadEnvAsInt("URL_SHORTENER_PORT"))
}
