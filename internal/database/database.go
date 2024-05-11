package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(url string) (*gorm.DB, error) {
	var tries int

	for {
		db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
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
