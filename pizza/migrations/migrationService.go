package migrations

import (
	"context"
	"database/sql"
)

type Service interface {
	CreateUserTable(tx *sql.Tx)
	CreatePizzaTable(tx *sql.Tx)
	CreateCartTable(tx *sql.Tx)
	CreateOrderTable(tx *sql.Tx)
	CreateOrderItemTable(tx *sql.Tx)
	CreateCartTotalTable(tx *sql.Tx)
	Run(ctx context.Context)
}
