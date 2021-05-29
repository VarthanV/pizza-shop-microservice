package message_queue

type OrderDetail struct {
	PizzaID  int `json:"pizza_id"`
	Quantity int `json:"quantity"`
}

type OrderQueueRequest struct {
	OrderUUID string        `json:"order_uuid"`
	Details   []OrderDetail `json:"details"`
}

type OrderUpdateStatusRequest struct {
	OrderUUID string `json:"order_uuid"`
	Status    string `json:"status"`
}
