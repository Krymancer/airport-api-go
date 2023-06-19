package models

import "gorm.io/gorm"

type Lugagge struct {
	gorm.Model
	Number int `json:"number"`
}
