package models

import "gorm.io/gorm"

type City struct {
	gorm.Model
	Name     string    `json:"name"`
	State    string    `json:"state"`
	Airports []Airport `json:"airports"`
}
