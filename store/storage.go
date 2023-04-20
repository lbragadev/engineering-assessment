package store

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/lbragadev/engineering-assessment/types"
)

type Storage interface {
	GetFoodTrucks() ([]*types.FoodTruck, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=elc_db password=postgres port=5003 sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) GetFoodTrucks() ([]*types.FoodTruck, error) {
	rows, err := s.db.Query("select * from food_trucks")
	if err != nil {
		return nil, err
	}

	foodTrucks := []*types.FoodTruck{}
	for rows.Next() {
		foodTruck, err := scanIntoFoodTruck(rows)
		if err != nil {
			return nil, err
		}
		foodTrucks = append(foodTrucks, foodTruck)
	}

	return foodTrucks, nil
}

func scanIntoFoodTruck(rows *sql.Rows) (*types.FoodTruck, error) {
	foodTruck := new(types.FoodTruck)
	err := rows.Scan(
		&foodTruck.ID,
		&foodTruck.LocationID,
		&foodTruck.Name,
		&foodTruck.Address,
		&foodTruck.Status,
		&foodTruck.FacilityType,
		&foodTruck.LocationDescription,
		&foodTruck.FoodItems,
		&foodTruck.Latitude,
		&foodTruck.Longitude,
		&foodTruck.CreatedAt)
	return foodTruck, err
}
