package models

import "gorm.io/gorm"

type Passenger struct {
	gorm.Model
	Name      string   `json:"name"`
	CPF       string   `json:"cpf"`
	BirthDate string   `json:"birth_date"`
	Email     string   `json:"email"`
	Tickests  []Ticket `json:"tickets"`
}
