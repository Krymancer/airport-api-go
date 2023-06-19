package models

type Ticket struct {
	Number int `json:"number"`
	SeatNumber string `json:"seat_number"`
	Lugagge []Lugagge `json:"lugagge"`
	Flight Flight `json:"flight"`
}