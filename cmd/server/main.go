package main

import (
	"fmt"
	"os"

	"github.com/marcos-nsantos/reservation-api/internal/database"
	"github.com/marcos-nsantos/reservation-api/internal/handler"
	"github.com/marcos-nsantos/reservation-api/internal/repository"
	"github.com/marcos-nsantos/reservation-api/internal/router"
	"github.com/marcos-nsantos/reservation-api/internal/service"
)

func main() {
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDB := os.Getenv("POSTGRES_DB")
	postgresPort := os.Getenv("POSTGRES_PORT")

	db, err := database.Connect(postgresHost, postgresUser, postgresPassword, postgresDB, postgresPort)
	if err != nil {
		fmt.Printf("failed to connect to database: %v\n", err)
		os.Exit(1)
	}

	if err = database.Migrate(db); err != nil {
		fmt.Printf("failed to migrate database: %v\n", err)
		os.Exit(1)
	}

	key := os.Getenv("SECRET_KEY")

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService, key)

	resourceRepo := repository.NewResourceRepository(db)
	resourceService := service.NewResourceService(resourceRepo)
	resourceHandler := handler.NewResourceHandler(resourceService)

	reservationRepo := repository.NewReservationRepository(db)
	reservationService := service.NewReservationService(reservationRepo, resourceRepo)
	reservationHandler := handler.NewReservationHandler(reservationService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	routerHandlers := router.Router{
		Key:                key,
		UserHandler:        userHandler,
		ResourceHandler:    resourceHandler,
		ReservationHandler: reservationHandler,
	}

	err = routerHandlers.SetupRouter().Run(":" + port)
	if err != nil {
		fmt.Printf("failed to start server: %v\n", err)
		os.Exit(1)
	}
}
