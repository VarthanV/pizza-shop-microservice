package repositories

import (
	"context"
	"database/sql"

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
