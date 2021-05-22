package repositories

import (
	"context"
	"database/sql"

	"github.com/VarthanV/pizza/pizza/models"
	"github.com/golang/glog"
)

type orderrepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) models.OrderRepository {

	return &orderrepository{
		db: db,
	}
}

func (o orderrepository) CreateOrder(ctx context.Context, order models.Order, userID string, cart *[]models.CartQueryResult) (err error) {
	/*
		1) Start a transaction -> TODO:To figure how to inject a trasnaction instead of sql.DB.
		2) Insert into orders table
		3) Convert all the cart items into order item
		4) Return success or err based on the outcome
		5) Start a go_routine parallel to make the cart_items inactive

	*/

	sql := `
	INSERT into orders(uuid,user_id,order_status) 
	values(?,?,?)
	
	`
	_, insertErr := o.db.ExecContext(ctx, sql, order.OrderUUID, userID, order.OrderStatus)
	if insertErr != nil {
		glog.Error("Error while inserting into orders table..", insertErr)

	}
	return insertErr
}

func (o orderrepository) GetOrderByUUID(ctx context.Context, uuid string) (*models.Order, error) {
	var order models.Order
	sql := `
	SELECT *
	FROM orders
	WHERE order_uuid = ?
	
	`
	err := o.db.QueryRowContext(ctx, sql, uuid).Scan(&order.ID, &order.OrderUUID, &order.UserID, &order.Total, &order.OrderStatus)
	if err != nil {
		glog.Errorf("Unable to query orders table %s ", err)
		return nil, err
	}
	return &order, nil
}

func (o orderrepository) GetOrdersByUserID(ctx context.Context, userId int) (*[]models.Order, error) {
	var orders []models.Order
	var order models.Order
	sql := `
	SELECT o.id
	FROM orders o
	WHERE user_id = ?
	`
	rows, err := o.db.QueryContext(ctx, sql, userId)
	if err != nil {
		glog.Errorf("Unable to query the order rows %s ", err)
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&order.ID, &order.OrderUUID, &order.UserID, &order.Total, &order.OrderStatus)
		if err != nil {
			glog.Errorf("Unable to scan the orders into struct 'Order' %s", err)
			return nil, err
		}
		orders = append(orders, order)
	}
	return &orders, nil
}
