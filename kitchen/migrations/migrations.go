package migrations

import (
	"context"
	"database/sql"
	"os"

	"github.com/golang/glog"
)

type migrations struct {
	db *sql.DB
}

func NewMigrationService(db *sql.DB) MigrationService {
	return &migrations{
		db: db,
	}
}

func (m migrations) migrateCookTable(ctx context.Context) {
	glog.Info("Creating Cooks table...")
	s := `
	CREATE TABLE IF NOT EXISTS cooks(
		id int not null AUTO_INCREMENT PRIMARY KEY,
		name varchar(200) not null,
		is_available int default 1
	  );
	`
	_, err := m.db.ExecContext(ctx, s)
	if err != nil {
		glog.Errorf("Unable to create cook table %s", err)
		os.Exit(-1)
	}
}
func (m migrations) migratePizzaCompletionTable(ctx context.Context) {
	glog.Info("Creating PizzaCompletionTable...")
	s := `
		CREATE TABLE IF NOT EXISTS pizza_completion(
			id int not null AUTO_INCREMENT PRIMARY KEY,
			pizza_id int not null,
			order_uuid int not null,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		)
	`
	_, err := m.db.ExecContext(ctx, s)
	if err != nil {
		glog.Error("Unable to create pizza process update table %s", err)
		os.Exit(-1)
	}
}

func (m migrations) migrateOrderCompletionTable(ctx context.Context) {
	glog.Info("Creating OrderCompletionTable")
	s := `
		CREATE TABLE IF NOT exists order_completion(
			id int not null AUTO_INCREMENT PRIMARY KEY,
			order_uuid varchar(200) not null,
			cook_id int not null,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		)	
	`
	_, err := m.db.ExecContext(ctx, s)
	if err != nil {
		glog.Error("Unable to create order process update table %s", err)
		os.Exit(-1)
	}
}

func (m migrations) RunMigrations(ctx context.Context) {
	m.migrateCookTable(ctx)
	m.migrateOrderCompletionTable(ctx)
	m.migratePizzaCompletionTable(ctx)
}
