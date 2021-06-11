package shared

import "time"

type SharedConstants struct {
	AccessTokenSecretKey  string
	RefreshTokenSecretKey string
}

type DBConnection struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

const (
	OrderStatusWaitingForCook = "Waiting for Cook"
	OrderStatusProcessing = "Processing"
	OrderStatusDelivered  = "Complete"
)

const (
	RedisKeyForOrders = "orders"
)

const (
	DeadlineForOrderSubmitRequest = 60 * time.Second
)
