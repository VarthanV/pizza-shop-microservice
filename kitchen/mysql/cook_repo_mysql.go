package mysql

import (
	"context"
	"database/sql"

	"github.com/VarthanV/kitchen/cooks"
	"github.com/VarthanV/kitchen/cooks/models"
	"github.com/golang/glog"
)

type cookrepomysql struct {
	db *sql.DB
}

func NewCookMysqlRepo(db *sql.DB) cooks.Repository {
	return cookrepomysql{
		db: db,
	}
}

func (c cookrepomysql) GetCookByID(ctx context.Context, id int) *models.Cook {
	var cook models.Cook
	s := `
		SELECT *
		FROM cooks
		WHERE id = ?
	`
	row := c.db.QueryRowContext(ctx, s, id)
	err := row.Scan(&cook.ID, &cook.Name, &cook.IsVegeterian, &cook.IsAvailbale)
	if err != nil {
		glog.Errorf("Error while scanning rows.. %s", err)
		return nil
	}
	return &cook
}

func (c cookrepomysql) GetAvailableCooks(ctx context.Context, IsVegeterian int) *[]models.Cook {
	var cook models.Cook
	var cooks []models.Cook

	s := `
		SELECT *
		FROM cooks 
		WHERE is_available =1
		AND is_vegeterian = ?

	`
	rows, err := c.db.QueryContext(ctx, s, IsVegeterian)
	if err != nil {
		glog.Error("Unable to query the available cooks %s", err)
		return nil
	}
	for rows.Next() {
		err = rows.Scan(&cook.ID, &cook.Name, &cook.IsVegeterian, &cook.IsAvailbale)
		if err != nil {
			glog.Errorf("Unable to scan the rows %s", err)
			return nil
		}
		cooks = append(cooks, cook)
	}
	return &cooks
}

func (c cookrepomysql) GetFirstAvailableCook(ctx context.Context, IsVegeterian int, cookChan chan *models.Cook) {
	var cook models.Cook
	s := `
		SELECT *
		FROM cooks
		WHERE is_available = 1
		AND is_vegeterian = ?
		LIMIT 1
	`
	go func() {
		row := c.db.QueryRowContext(ctx, s, IsVegeterian)
		err := row.Scan(&cook.ID, &cook.Name, &cook.IsVegeterian, &cook.IsAvailbale)
		if err != nil {
			glog.Errorf("Error while scanning rows.. %s", err)
			cookChan <- nil

		} else {
			cookChan <- &cook
		}
	}()
}
