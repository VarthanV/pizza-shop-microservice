package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/VarthanV/pizza/pizza/models"
	"github.com/golang/glog"
)

type orderitemrepository struct {
	db *sql.DB
}

func NewOrderItemRepository(db *sql.DB) models.OrderItemRepository {
	return &orderitemrepository{
		db: db,
	}
}

func (o orderitemrepository) GetOrderItemByID(ctx context.Context, id int) (*models.OrderItem, error) {
	var orderItem models.OrderItem
	s := `
	SELECT * 
	FROM order_item
	WHERE id = ?
`
	row := o.db.QueryRowContext(ctx, s, id)
	if row == nil {
		glog.Errorf("No order item found for the given ID %d", id)
		return nil, errors.New("no order item found for the given id")
	}
	err := row.Scan(&orderItem.ID, &orderItem.OrderID, &orderItem.PizzaID, &orderItem.Quantity, &orderItem.Price)
	if err != nil {
		glog.Error("Unable to scan the row..")
		return nil, err
	}
	return &orderItem, nil
}

func (o orderitemrepository) GetOrderItemsByOrderID(ctx context.Context, orderID int) (*[]models.OrderItem, error) {

	panic("implement me")
}

func (o orderitemrepository) AddOrderItem(ctx context.Context, pizzaID int, orderUUID string, quantity int, price int) error {
	s := `
	INSERT INTO order_item
	(order_uuid,pizza_id,price,quantity)
	values(?,?,?,?)
	`
	_, err := o.db.ExecContext(ctx, s, orderUUID, pizzaID, price, quantity)
	if err != nil {
		glog.Errorf("Unable to create order item %s", err)
		return err
	}
	return nil
}
