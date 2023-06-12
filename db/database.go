package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env")
	}
	d, err := sql.Open("pgx", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	DB = d
}
