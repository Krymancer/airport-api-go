package models

import "gorm.io/gorm"

type Flight struct {
	gorm.Model
	Number        int           `json:"number"`
	Departure     string        `json:"departure"`
	Arrival       string        `json:"arrival"`
	OriginID      uint          `json:"origin_id"`
	DestinationID uint          `json:"destination_id"`
	FlightClasses []FlightClass `json:"flight_class"`
	Tickets       []Ticket      `json:"tickets"`
}
