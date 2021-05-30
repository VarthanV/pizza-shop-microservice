package models

type Cook struct {
	ID           int    `json:"int"`
	Name         string `json:"name"`
	IsVegeterian int    `json:"is_vegeterian"`
	IsAvailbale  int    `json:"is_available"`
}
