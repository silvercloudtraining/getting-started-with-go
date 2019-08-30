package main

import (
	"log"
	"net/http"

	"github.com/silvercloudtraining/coffeeservice/coffee"

	"github.com/silvercloudtraining/coffeeservice/routing"
)

const port = ":3000"

func main() {
	inv := coffee.NewInventory()
	routing.RegisterRoutes(inv)
	var err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
