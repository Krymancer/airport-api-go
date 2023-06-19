package models

type Passenger struct {
	Name string `json:"name"`
	CPF string `json:"cpf"`
	BirthDate string `json:"birth_date"`
	Email string `json:"email"`
}