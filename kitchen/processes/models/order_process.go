package models

type OrderProcess struct {
	ID                 int    `json:"id,omitempty"`
	OrderUUID          string `json:"order_uuid"`
	CookID             int    `json:"cook_id"`
	TimeTakenInSeconds int    `json:"time_taken_in_seconds"`
}
