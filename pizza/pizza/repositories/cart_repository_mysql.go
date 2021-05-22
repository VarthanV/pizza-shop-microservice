package repositories

import (
	"context"
	"database/sql"

	"github.com/VarthanV/pizza/pizza/models"
	"github.com/golang/glog"
)

type cartrepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) models.CartRepository {
	return &cartrepository{
		db: db,
	}
}

func (c cartrepository) AddItem(ctx context.Context, itemId int, userId string, quantity int, price int) error {

	s := `
	INSERT into
  cart (user_id, pizza_id, quantity, price,is_active)
values(?,?,?,?,?)

	`
	_, insertErr := c.db.ExecContext(ctx, s, userId, itemId, quantity, price, 1)
	if insertErr != nil {
		glog.Errorf("Unable to insert into the cart table ... %s", insertErr)
		return insertErr

	}

	return nil
}

func (c cartrepository) EditItem(ctx context.Context, cartItemId int, itemId int, quantity int, price int, userId string) error {
	s := `
	UPDATE cart
	set pizza_id = ?,
	quantity = ?,
	price = ?
	WHERE id = ?
	AND user_id = ?
`
	_, insertErr := c.db.ExecContext(ctx, s, itemId, quantity, price, cartItemId, userId)
	if insertErr != nil {
		glog.Errorf("Unable to update cart with ID %d  %s", cartItemId, insertErr)
	}
	return nil
}

func (c cartrepository) DeleteItem(ctx context.Context, cartItemId int, userId string) error {
	s := `
	DELETE from cart
	where id = ?
	and user_id = ?
	`
	_, insertErr := c.db.ExecContext(ctx, s, cartItemId, userId)
	if insertErr != nil {
		glog.Errorf("Unable to delete cart item... %s", insertErr)
	}
	return nil
}

func (c cartrepository) GetCart(ctx context.Context, userId string) (*[]models.CartQueryResult, error) {
	var carts []models.CartQueryResult
	var cart models.CartQueryResult

	s := `
	SELECT
	c.id,
  	p.name,
	c.price,
	c.quantity,
	p.id,
  	p.is_vegeterian
	FROM
  	cart AS c
  	INNER join 
	pizzas p on p.id = c.pizza_id
	where
  	c.user_id = ?
	AND 
	c.is_active = 1
	`
	rows, err := c.db.QueryContext(ctx, s, userId)
	if err != nil {
		glog.Errorf("Unable to query the order rows %s", err)
	}
	for rows.Next() {
		err := rows.Scan(&cart.ID, &cart.PizzaName, &cart.Price, &cart.Quantity, &cart.PizzaID, &cart.IsVegeterian)
		if err != nil {
			glog.Errorf("Unable to scan rows for the cart model %s", err)
		}
		carts = append(carts, cart)
	}
	return &carts, nil

}

func (c cartrepository) GetCartItem(ctx context.Context, itemId int, userId string) *models.Cart {
	//TODO: Make the query more presentable
	var cart models.Cart
	s := `
		SELECT c.id
		FROM cart c
		WHERE pizza_id = ? AND user_id= ?
	`
	row := c.db.QueryRowContext(ctx, s, itemId, userId)
	err := row.Scan(&cart.ID)
	if err != nil {
		glog.Errorf("Unable to scan row %s", err)
		return nil
	}
	return &cart
}

func (c cartrepository) MakeItemInactive(ctx context.Context, itemId int) error {
	s := `
	UPDATE cart
	SET    
	is_active = 0
	WHERE id = ?
	`
	_, err := c.db.ExecContext(ctx, s, itemId)
	if err != nil {
		glog.Errorf("Error updating cart %s ", err)
	}
	return err
}
