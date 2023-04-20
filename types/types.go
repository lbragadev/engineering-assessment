package types

import (
	"time"
)

type FoodTruck struct {
	ID                  int       `json:"id"`
	LocationID          int64     `json:"locationId"`
	Name                string    `json:"name"`
	Address             string    `json:"address"`
	Status              string    `json:"status"`
	FacilityType        string    `json:"facilityType"`
	LocationDescription string    `json:"locationDescription"`
	FoodItems           string    `json:"foodItems"`
	Latitude            float64   `json:"latitude"`
	Longitude           float64   `json:"longitude"`
	CreatedAt           time.Time `json:"createdAt"`
}
