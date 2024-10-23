package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func CreateConnection() *sql.DB {

	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	err = db.Ping()

	if err != nil {
		fmt.Println("Database is NOK")
		log.Fatalln(err)
	}

	return db
}
