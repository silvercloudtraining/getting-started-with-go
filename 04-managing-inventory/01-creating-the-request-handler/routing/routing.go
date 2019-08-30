package routing

import "net/http"

func RegisterRoutes() {
	http.HandleFunc("/users", userHandler)
	http.HandleFunc("/users/", userHandler)

	ch := &coffeeHandler{}
	http.Handle("/coffee", ch)
	http.Handle("/coffee/", ch)
}
