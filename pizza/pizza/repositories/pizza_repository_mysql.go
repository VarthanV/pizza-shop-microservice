package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/VarthanV/pizza/pizza/models"
	"github.com/golang/glog"
)

type repository struct {
	db *sql.DB
}

func NewPizzaMysqlRepository(db *sql.DB) models.PizzaRepository {
	return &repository{
		db: db,
	}
}

func (r repository) GetPizzaByID(ctx context.Context, id int) (resultPizza models.Pizza, err error) {
	var pizza models.Pizza
	query := `
	SELECT * 
	from pizzas
	WHERE id = ?
	`
	row := r.db.QueryRowContext(ctx, query, id)
	if row == nil {
		glog.Errorf("Unable to get pizza with id %d %s", id, err)
		return models.Pizza{}, errors.New("no pizza found ")
	}
	err = row.Scan(&pizza.ID, &pizza.Name, &pizza.Price, &pizza.IsVegeterian)
	if err != nil {
		glog.Errorf("Unable to scan the rows")
		return models.Pizza{}, err
	}
	return pizza, nil
}

func (r repository) GetAllPizzas(ctx context.Context, isVegetarian int) (pizzas []models.Pizza, err error) {
	var queriedPizzas []models.Pizza
	var pizza models.Pizza

	sql := `
	SELECT * 
	from pizzas
	WHERE is_vegeterian = ?
	
	`
	rows, err := r.db.Query(sql, isVegetarian)
	if err != nil {
		glog.Errorf("Unable to get pizzas  %f", err)
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&pizza.ID, &pizza.Name, &pizza.Price, &pizza.IsVegeterian)
		queriedPizzas = append(queriedPizzas, pizza)
		if err != nil {
			glog.Error("Unable to scan pizzas...", err)
			return nil, err
		}
	}
	return queriedPizzas, nil
}
