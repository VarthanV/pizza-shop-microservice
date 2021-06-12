package repositories

import (
	"context"
	"database/sql"

	"github.com/VarthanV/pizza/pizza"
	"github.com/golang/glog"
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
	SET order_status = ?
	WHERE uuid = ?	
	`
	_, err := our.db.ExecContext(ctx, s, status, orderUUID)
	if err != nil {
		glog.Errorf("Error is %s",err)
	}
	return err
}
