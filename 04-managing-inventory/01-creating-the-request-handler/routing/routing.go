package routing

import "net/http"

func RegisterRoutes() {
	http.HandleFunc("/users", userHandler)
	http.HandleFunc("/users/", userHandler)

	ih := &inventoryHandler{}
	http.Handle("/inventory", ih)
	http.Handle("/inventory/", ih)
}
