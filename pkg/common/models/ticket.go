package models

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	Number     int       `json:"number"`
	SeatNumber string    `json:"seat_number"`
	Lugagge    []Lugagge `json:"lugagge"`
	Flight     Flight    `json:"flight"`
	Passenger  Passenger `json:"passenger"`
}
