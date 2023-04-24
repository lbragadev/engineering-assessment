package main

import (
	"log"

	"github.com/lbragadev/engineering-assessment/api"
	"github.com/lbragadev/engineering-assessment/store"
)

func main() {
	//Add filtering for food_items field
	store, err := store.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	server := api.NewAPIServer("0.0.0.0:8080", store)
	server.Run()

}
