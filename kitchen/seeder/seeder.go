package seeder

import (
	"database/sql"

	"github.com/VarthanV/kitchen/cooks/models"
	"github.com/golang/glog"
)

type seederservice struct {
	db *sql.DB
}

func NewSeederService(db *sql.DB) Service {
	return &seederservice{
		db: db,
	}
}

func (ss seederservice) seedCooksData() {
	glog.Info("Seeding data for cook table...")
	cooks := []models.Cook{
		{Name: "Walter white"},
		{Name: "Gustavo Fring"},
		{Name: "Jessie Pinkman"},
		{Name: "Hank Shrader"},
		{Name: "Saul Goodman"},
	}

	query := `
	
	INSERT into cooks (name)
	values (?)
	
	`
	for _, cook := range cooks {
		_, err := ss.db.Exec(query, cook.Name)
		if err != nil {
			glog.Errorf("Unable to seed data...", err)
		}
	}
}

func (ss seederservice) SeedData() {
	ss.seedCooksData()

}
