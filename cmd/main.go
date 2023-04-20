package main

import (
	"log"

	"github.com/lbragadev/engineering-assessment/api"
	"github.com/lbragadev/engineering-assessment/store"
)

func main() {

	store, err := store.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	server := api.NewAPIServer(":3000", store)
	server.Run()
}
