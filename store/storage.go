package store

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/lbragadev/engineering-assessment/types"
)

type Storage interface {
	GetFoodTrucks() ([]*types.FoodTruck, error)
}

type PostgresStore struct {
	db *sql.DB
}

func Init() {
	err := godotenv.Load("prod.env")
	if err != nil {
		log.Fatal(".env file couldn't be loaded")
	}
	EnvVars = loadEnvVars()
}

func NewPostgresStore() (*PostgresStore, error) {
	Init()
	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s port=%s sslmode=disable",
		EnvVars.DbUser,
		EnvVars.DbPass,
		EnvVars.DbName,
		EnvVars.DbPort,
	)
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

func generateGoogleMapsUrl(lat float64, long float64) string {
	latStr := fmt.Sprintf("%f", lat)
	longStr := fmt.Sprintf("%f", long)
	res := fmt.Sprintf("http://maps.google.com/maps?z=12&t=m&q=loc:%s+%s", latStr, longStr)
	return res
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
	foodTruck.GoogleMapsUrl = generateGoogleMapsUrl(foodTruck.Latitude, foodTruck.Longitude)
	return foodTruck, err
}
