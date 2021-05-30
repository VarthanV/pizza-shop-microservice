package mysql

import (
	"context"
	"database/sql"

	"github.com/VarthanV/kitchen/processes"
)

type processOrderRepo struct {
	db *sql.DB
}

func NewProcessOrderRepository(db *sql.DB) processes.OrderProcessRepository {
	return &processOrderRepo{
		db: db,
	}
}

func (pr processOrderRepo) CompleteOrder(ctx context.Context, orderID int, cookID int) error {

	return nil
}
