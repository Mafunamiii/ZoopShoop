package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func ConnectToDatabase() (*sql.DB, error) {
	log.Info().Msg("Connecting to the database")

	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	//dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require",
		dbHost, dbUser, dbPassword, dbName)

	log.Info().Msgf("Connecting to the database with connection string: %s", connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to database")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Error().Err(err).Msg("Database ping failed")
		return nil, err
	}

	log.Info().Msg("Successfully connected to the database")
	return db, nil
}
