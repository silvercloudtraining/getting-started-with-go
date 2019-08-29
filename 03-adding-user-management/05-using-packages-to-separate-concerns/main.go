package main

import (
	"log"
	"net/http"

	"github.com/silvercloudtraining/coffeeservice/routing"
)

const port = ":3000"

func main() {
	routing.RegisterRoutes()
	var err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
