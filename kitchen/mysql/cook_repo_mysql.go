package mysql

import (
	"context"
	"database/sql"

	"github.com/VarthanV/kitchen/cooks"
	"github.com/VarthanV/kitchen/cooks/models"
)

type cookrepomysql struct{
	db *sql.DB
}

func NewCookMysqlRepo(db *sql.DB) cooks.Repository{
	return cookrepomysql{
		db: db,
	}
}

func (c cookrepomysql) GetCookByID(ctx context.Context,id int) *models.Cook{
	
	return nil
}