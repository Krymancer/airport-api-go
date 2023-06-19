package models

import "gorm.io/gorm"

type Flight struct {
	gorm.Model
	Number        int           `json:"number"`
	Departure     string        `json:"departure"`
	Arrival       string        `json:"arrival"`
	Origin        Airport       `json:"origin"`
	Destination   Airport       `json:"destination"`
	FlightClasses []FlightClass `json:"flight_class"`
	Tickets       []Ticket      `json:"tickets"`
}
