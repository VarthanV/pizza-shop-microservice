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

func (m migrations) RunMigrations(ctx context.Context) {
	m.migrateCookTable(ctx)
}
