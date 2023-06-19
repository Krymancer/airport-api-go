package db

import (
	"log"

	"github.com/krymancer/airport-api-go/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(
		&models.Airport{},
		&models.City{},
		&models.FlightClass{},
		&models.Flight{},
		&models.Lugagge{},
		&models.Manager{},
		&models.Passenger{},
		&models.Ticket{},
		&models.Visitor{},
	)

	return db
}
