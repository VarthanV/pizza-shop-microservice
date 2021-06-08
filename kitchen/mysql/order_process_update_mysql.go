package mysql

import (
	"context"
	"database/sql"

	"github.com/VarthanV/kitchen/processes"
)

type orderprocessupdatemysql struct {
	db *sql.DB
}

var err error
var s string

func NewOrderProcessUpdateRepoMysql(db *sql.DB) processes.OrderProcessUpdateRepo {

	return &orderprocessupdatemysql{
		db: db,
	}
}

func (o orderprocessupdatemysql) UpdateOrderProcces(ctx context.Context, orderUUID string, cookID int) error {
	s = `
		INSERT into order_processing_updates (order_uuid , cook_id)
		VALUES (?,?)
	`
	_, err = o.db.ExecContext(ctx, s, orderUUID, cookID)
	return err
}

func (o orderprocessupdatemysql) UpdateOrderItemProcess(ctx context.Context, pizzaID int, orderUUID string) error {
	s := `
	INSERT into pizza_processing_updates (pizza_id,order_uuid)
	VALUES (?,?)
	
	`
	_, err = o.db.ExecContext(ctx, s, pizzaID, orderUUID)
	return err
}
