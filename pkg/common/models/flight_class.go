package models

import "gorm.io/gorm"

type FlightClass struct {
	gorm.Model
	Name          string  `json:"name"`
	Type          int     `json:"type"`
	NumberOfSeats int     `json:"number_of_seats"`
	PricePerSeat  float64 `json:"price_per_seat"`
}
