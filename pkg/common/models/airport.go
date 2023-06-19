package models

import "gorm.io/gorm"

type Airport struct {
	gorm.Model
	Name     string `json:"name"`
	IATAcode string `json:"iata_code"`
	City     City   `json:"city"`
}
