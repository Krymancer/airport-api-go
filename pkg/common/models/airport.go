package models

import "gorm.io/gorm"

type Airport struct {
	gorm.Model
	Name     string `json:"name"`
	IATACode string `json:"iata_code"`
	CityID   uint   `json:"city_id"`
}
