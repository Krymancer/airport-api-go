package models

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	Number      int       `json:"number"`
	SeatNumber  string    `json:"seat_number"`
	Lugagge     []Lugagge `json:"lugagge"`
	FlightID    uint      `json:"flight"`
	PassengerID uint      `json:"passenger_id"`
	VisitorID   uint      `json:"visitor_id"`
}
