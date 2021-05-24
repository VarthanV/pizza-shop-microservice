package models

type PizzaProcess struct {
	ID                 int `json:"id,omitempty"`
	PizzaID            int `json:"pizza_id"`
	CookID             int `json:"cook_id"`
	TimeTakenInSeconds int `json:"time_taken_in_seconds"`
}
