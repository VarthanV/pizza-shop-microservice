package migrations

import (
	"context"
	"database/sql"
	"github.com/golang/glog"
	"os"
)

type migrationservice struct {
	db  *sql.DB
	ctx context.Context
}

/*
	Try to create all the tables that is necessary for our app
	to function , If anyone of it fails exit gracefully
	and leave a error message.
*/
func handleError(tx *sql.Tx) {
	err := tx.Rollback()
	if err != nil {
		glog.Errorf("Unable to rollback the transaction..")
		os.Exit(-1)

	}

}

func (m migrationservice) CreateUserTable(tx *sql.Tx) {
	query := `
  CREATE table if not exists users (
  id varchar(200) not null PRIMARY KEY,
  name varchar(200) not null,
  email varchar(200) not null,
  password varchar(200) not null,
  phone_number varchar(10) not null
)
`
	glog.Info("Creating Table Users..")
	_, err := tx.ExecContext(m.ctx, query)
	if err != nil {
		glog.Errorf("Unable to create table Users %s", err)
		handleError(tx)
	}
}

func (m migrationservice) CreatePizzaTable(tx *sql.Tx) {
	query := `
	CREATE table if not exists pizzas(
  id int not null AUTO_INCREMENT PRIMARY KEY ,
  name varchar(200) not null,
  price int not null default 100,
  is_vegeterian int not null default 0
);
`
	glog.Info("Creating Table Pizzas..")
	_, err := tx.ExecContext(m.ctx, query)
	if err != nil {
		glog.Errorf("Unable to create table Pizzas %s", err)
		handleError(tx)
	}
}

func (m migrationservice) CreateCartTable(tx *sql.Tx) {
	query := `
  CREATE table if not exists cart(
  id int not null AUTO_INCREMENT primary key ,
  pizza_id int not null ,
  quantity int not null default 1,
  price int not null ,
  user_id varchar(200) not null ,
  is_active int not null default 1,
  FOREIGN KEY(user_id) references users(id),
  FOREIGN KEY (pizza_id) references pizzas(id)
);`
	glog.Info("Creating Table Cart..")
	_, err := tx.ExecContext(m.ctx, query)
	if err != nil {
		glog.Errorf("Unable tp create table Cart %s", err)
		handleError(tx)
	}
}

func (m migrationservice) CreateOrderTable(tx *sql.Tx) {
	query := `
	CREATE table if not exists orders(
	id int not null AUTO_INCREMENT PRIMARY KEY ,
	uuid varchar(200) unique not null,
	user_id varchar(200) not null,
	order_status varchar(200) default 'Placed',
	FOREIGN KEY(user_id) references  users(id)                               
)
`
	glog.Info("Creating Table Order...")
	_, err := tx.ExecContext(m.ctx, query)
	if err != nil {
		glog.Errorf("Unable to create table Order %s", err)
		handleError(tx)
	}
}

func (m migrationservice) CreateOrderItemTable(tx *sql.Tx) {
	query := `
	CREATE table if not exists order_item(
	id int not null AUTO_INCREMENT primary key ,
	order_uuid varchar(200) not null ,
	pizza_id int not null ,
	price int not null ,
	quantity int not null ,
	FOREIGN KEY (order_uuid) references  orders(uuid) ON DELETE  CASCADE ,
	FOREIGN KEY(pizza_id) references pizzas(id)
)
`
	glog.Info("Creating Table OrderItem...")
	_, err := tx.ExecContext(m.ctx, query)
	if err != nil {
		glog.Errorf("Unable to create table OrderItem %s", err)
		handleError(tx)
	}

}

func (m migrationservice) CreateCartTotalTable(tx *sql.Tx) {
	query := `
	CREATE table if not exists cart_total(
	cart_id int not null,
	total int,
	FOREIGN KEY(cart_id) references cart(id)
)
`
	glog.Info("Creating Table CartTotal...")
	_, err := tx.ExecContext(m.ctx, query)
	if err != nil {
		glog.Errorf("Unable to create table CartTotal %s", err)
		handleError(tx)
	}
}

func (m migrationservice) Run(ctx context.Context) {
	tx, _ := m.db.BeginTx(ctx, nil)

	m.CreateUserTable(tx)
	m.CreatePizzaTable(tx)
	m.CreateCartTable(tx)
	m.CreateOrderTable(tx)
	m.CreateOrderItemTable(tx)
	m.CreateCartTotalTable(tx)
	tx.Commit()
}

func NewMigrationService(db *sql.DB, ctx context.Context) Service {
	return &migrationservice{
		db:  db,
		ctx: ctx,
	}
}
