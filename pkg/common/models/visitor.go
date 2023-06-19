package models

type Visitor struct {
	Name string `json:"name"`
	CPF string `json:"cpf"`
	BirthDate string `json:"birth_date"`
	Email string `json:"email"`
	Tickets []Ticket `json:"tickets"`
}