package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/marcos-nsantos/reservation-api/internal/entity"
)

func Connect(host, user, password, dbname, port string) (*gorm.DB, error) {
	var tries int

	databaseURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		dbname,
		port,
	)

	for {
		db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})

		if err != nil {
			if tries >= 5 {
				return nil, fmt.Errorf("failed to connect to database: %w", err)
			}

			fmt.Println("database is not ready yet, retrying in 5 seconds")
			time.Sleep(5 * time.Second)

			tries++
		} else {
			return db, nil
		}
	}
}

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&entity.User{}); err != nil {
		return fmt.Errorf("failed to migrate user: %w", err)
	}

	if err := db.AutoMigrate(&entity.Resource{}); err != nil {
		return fmt.Errorf("failed to migrate resource: %w", err)
	}

	if err := db.AutoMigrate(&entity.Reservation{}); err != nil {
		return fmt.Errorf("failed to migrate reservation: %w", err)
	}

	return nil
}
