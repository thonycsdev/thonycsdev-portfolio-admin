package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"thonycsdev/thonycsdev-portfolio-admin/domain"

	_ "github.com/lib/pq"
)

func CreateConnection() {
	//Cria uma conexao com o banco de dados

	connStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	result, err := db.Query("SELECT * from Experiences")
	if err != nil {
		log.Fatal(err)
	}

	for result.Next() {
		var experience domain.Experiences

		err = result.Scan(&experience.Id, &experience.Name, &experience.Level, &experience.Description, &experience.StartedYear, &experience.EndedYear)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(experience.Name)
	}

	defer db.Close()
}
