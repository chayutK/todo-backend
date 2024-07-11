package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Sync() {
	DB = database()
}

func database() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URI"))

	if err != nil {
		log.Fatal("Connection to database error", err)
	}

	fmt.Println("Database sync successfully")

	return db
}
