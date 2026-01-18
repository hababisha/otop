package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	if err := godotenv.Load(); err != nil{
		log.Fatal("error loading .env")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connstr := fmt.Sprintf("sslmode=disable host=%s port=%s user=%s password=%s dbname=%s", host, port, username, password, dbName)
	var err error
	DB, err = sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal("failed to open db: ", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("failed to connect to db: ", err)
	}

}