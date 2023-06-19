package models

type Airport struct {
	Name string `json:"name"`
	IATAcode string `json:"iata_code"`
	city City `json:"city"`
}