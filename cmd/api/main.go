package main

import (
	"database/sql"
	"log"
	"rest-apievent/internal/database"
	"rest-apievent/internal/env"
)

type aplication struct {
	port      int
	jwtSecret string
	models    *database.Models
}

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	models := database.NewModels(db)
	app := &aplication{
		port:      env.GetEnvInt("PORT", 8080),
		jwtSecret: env.GetEnvString("JWT_SECRET", "secret"),
		models:    models,
	}

	if err := app.server(); err != nil {
		log.Fatal(err)
	}
}
