package repositories

import (
	"context"
	"database/sql"

	"github.com/VarthanV/pizza/pizza"
)

type orderupdaterepository struct {
	db *sql.DB
}

func NewOrderUpdateRepository(db *sql.DB) pizza.OrderUpdateRepo {
	return &orderupdaterepository{
		db: db,
	}
}

func (our orderupdaterepository) UpdateOrderStatus(ctx context.Context, orderUUID string, status string) error {
	s := `
	UPDATE orders
	SET status = ?
	WHERE uuid = ?	
	`
	_, err := our.db.ExecContext(ctx, s, status, orderUUID)
	return err
}
