package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func DatabaseSync() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URI"))

	if err != nil {
		log.Fatal("Connection to database error", err)
	}

	fmt.Println("Database sync successfully")

	return db
}
