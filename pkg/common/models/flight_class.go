package models

type FlightClass struct {
	Name string `json:"name"`
	Type int `json:"type"`
	numberOfSeats int `json:"number_of_seats"`
	pricePerSeat float64 `json:"price_per_seat"`
}