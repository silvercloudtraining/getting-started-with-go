package routing

import "net/http"

type inventoryHandler struct{}

func (ih inventoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Inventory handler reporting for duty!"))
}
