package routing

import "net/http"

type coffeeHandler struct{}

func (ch coffeeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Coffee handler reporting for duty!"))
}
