package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

func CreateConnection() (*sqlx.DB, error) {
	// Get database details from environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	DBName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	return sqlx.Connect("postgres", fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s",
		host, user, DBName, password,
	))
	// https://github.com/volatiletech/sqlboiler
}
