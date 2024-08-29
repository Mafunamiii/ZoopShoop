package database

import (
	"database/sql"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectToDatabase() (*sql.DB, error) {
	log.Info().Msg("Connecting to the database")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		dbHost, dbPort, dbName, dbUser, dbPassword)

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
