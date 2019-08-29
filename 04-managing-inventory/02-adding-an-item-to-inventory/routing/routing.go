package routing

import (
	"net/http"

	"github.com/silvercloudtraining/coffeeservice/inventory"
)

func RegisterRoutes(inv *inventory.Inventory) {
	http.HandleFunc("/users", userHandler)
	http.HandleFunc("/users/", userHandler)

	ih := &inventoryHandler{inventory: inv}
	http.Handle("/inventory", ih)
	http.Handle("/inventory/", ih)
}
