package handlers

type AddToCartRequest struct {
	PizzaID  int `json:"pizza_id"`
	Quantity int `json:"quantity"`
	Price    int `json:"price"`
}
