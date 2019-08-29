package main

import (
	"log"
	"net/http"

	"github.com/silvercloudtraining/coffeeservice/routing"
)

func main() {
	routing.RegisterRoutes()
	var err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
