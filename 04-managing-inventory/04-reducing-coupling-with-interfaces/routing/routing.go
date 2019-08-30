package routing

import (
	"net/http"
)

func RegisterRoutes(inv coffeeInventory) {
	http.HandleFunc("/users", userHandler)
	http.HandleFunc("/users/", userHandler)

	ch := NewCoffeeHandler(inv)
	http.Handle("/coffee", ch)
	http.Handle("/coffee/", ch)
}
