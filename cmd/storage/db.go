package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err = sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort))

	if err != nil {
		panic(err.Error())
	}
	//InitializeData()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	InitializeData()
	fmt.Println("Successfully connected to database")
}

func GetDB() *sql.DB {
	return db
}

func InitializeData() {

	path := filepath.Join("/app/sql/products_init.sql")

	c, ioErr := os.ReadFile(path)
	if ioErr != nil {
		panic(ioErr.Error())
	}

	_, err := db.Exec(string(c))
	if err != nil {
		panic(err.Error())
	}
}
