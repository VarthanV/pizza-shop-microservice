package shared

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
	OrderStatusCreated    = "order-created"
	OrderStatusProcessing = "order-processing"
	OrderStatusDelivered  = "order-delivered"
)
