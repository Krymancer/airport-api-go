package models

import "gorm.io/gorm"

type Manager struct {
	gorm.Model
	Name      string `json:"name"`
	CPF       string `json:"cpf"`
	BirthDate string `json:"birth_date"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Username  string `json:"username"`
}
