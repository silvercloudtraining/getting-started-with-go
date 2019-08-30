package routing

import (
	"net/http"

	"github.com/silvercloudtraining/coffeeservice/coffee"
)

func RegisterRoutes(inv *coffee.Inventory) {
	http.HandleFunc("/users", userHandler)
	http.HandleFunc("/users/", userHandler)

	ch := NewCoffeeHandler(inv)
	http.Handle("/coffee", ch)
	http.Handle("/coffee/", ch)
}
