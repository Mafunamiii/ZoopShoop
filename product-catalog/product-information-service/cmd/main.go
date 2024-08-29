package main

import (
	"github.com/Mafunamiii/ZoopShoop/tree/main/product-catalog/product-information-service/pkg/database"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Starting the product information service")
	db, err := database.ConnectToDatabase()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to the database")
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Error().Err(err).Msg("Error closing database connection")
		}
	}()
}
