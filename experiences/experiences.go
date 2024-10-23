package experiences

import (
	"fmt"
	"log"
	"thonycsdev/thonycsdev-portfolio-admin/postgres"
)

type Experience struct {
	Id                       int
	Name, Level, Description string
	StartedYear, EndedYear   string
}

// Cria uma novo instancia de Experiencias
func New(name, level, description, started_year, ended_year string) *Experience {
	exp := Experience{
		Name:        name,
		Level:       level,
		Description: description,
		StartedYear: started_year,
		EndedYear:   ended_year,
	}
	return &exp
}

// Busca todas as experiencias do banco de dados
func GetExperiences() *[]Experience {
	database := postgres.CreateConnection()
	defer database.Close()

	query := "SELECT * FROM Experiences;"
	results := []Experience{}
	rows, _ := database.Query(query)
	for rows.Next() {
		var experience Experience
		err := rows.Scan(&experience.Id, &experience.Name, &experience.Description, &experience.Level, &experience.StartedYear, &experience.EndedYear)
		if err != nil {
			log.Fatalln(err)
		}
		results = append(results, experience)
	}

	return &results
}

// Adiciona uma experiencia no banco de dados
func AddExperience(experience *Experience) {
	database := postgres.CreateConnection()
	defer database.Close()

	query := "INSERT INTO Experiences (name, level, description, started_year, ended_year) VALUES ($1,$2,$3,$4,$5)"
	result, err := database.Exec(query, experience.Name, experience.Level, experience.Description, experience.StartedYear, experience.EndedYear)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
