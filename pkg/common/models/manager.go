package models

type Manager struct {
	Name string `json:"name"`
	CPF string `json:"cpf"`
	BirthDate string `json:"birth_date"`
	Email string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}