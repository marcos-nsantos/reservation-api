package main

import (
	"fmt"
	"os"

	"github.com/marcos-nsantos/reservation-api/internal/database"
	"github.com/marcos-nsantos/reservation-api/internal/router"
)

func main() {
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDB := os.Getenv("POSTGRES_DB")
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresPort := os.Getenv("POSTGRES_PORT")

	connectionURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		postgresHost,
		postgresUser,
		postgresPassword,
		postgresDB,
		postgresPort,
	)

	fmt.Println("Connecting to database...")
	db, err := database.Connect(connectionURL)
	if err != nil {
		fmt.Printf("failed to connect to database: %v\n", err)
		os.Exit(1)
	}

	if err = database.Migrate(db); err != nil {
		fmt.Printf("failed to migrate database: %v\n", err)
		os.Exit(1)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := router.SetupRouter()
	r.Run(":" + port)
}
