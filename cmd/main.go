package main

import (
	"log"

	"github.com/lbragadev/engineering-assessment/api"
	"github.com/lbragadev/engineering-assessment/store"
)

func main() {
	//TODO:
	//Add unit tests
	//Add filtering for food_items field
	//Add google_maps_url_field
	//DONE

	//Deployment
	//deploy on aws
	//dockerize this app
	//DONE
	//use env vars for db creds
	//DONE
	store, err := store.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	server := api.NewAPIServer("0.0.0.0:8080", store)
	server.Run()

}
