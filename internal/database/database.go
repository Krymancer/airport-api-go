package database

import (
	"log"

	"github.com/krymancer/airport-api-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(
		&models.Manager{},
		&models.City{},
		&models.Airport{},
		&models.Flight{},
		&models.FlightClass{},
		&models.Passenger{},
		&models.Visitor{},
		&models.Ticket{},
		&models.Lugagge{},
	)

	return db
}
