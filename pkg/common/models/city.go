package models

type City struct {
	Name string `json:"name"`
	State string `json:"state"`
	Airports []Airport `json:"airports"`
}