package main

import (
	"github.com/silvercloudtraining/coffeeservice/inventory"
	"log"
	"net/http"

	"github.com/silvercloudtraining/coffeeservice/routing"
)

const port = ":3000"

func main() {
	inv := inventory.Inventory{}
	routing.RegisterRoutes(&inv)
	var err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
