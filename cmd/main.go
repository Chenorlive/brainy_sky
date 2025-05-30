package main

import (
	"log"
	"os"

	"github.com/Chenorlive/brainy/cmd/api"
	db "github.com/Chenorlive/brainy/database"
)

const defaultPort = "8080"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Initialize the database connection
	db, err := db.SetupDB(true)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the server
	server := api.NewServer(":"+port, db)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
