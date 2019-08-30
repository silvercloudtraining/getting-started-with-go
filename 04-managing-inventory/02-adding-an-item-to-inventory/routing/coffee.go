package routing

import (
	"encoding/json"
	"net/http"

	"github.com/silvercloudtraining/coffeeservice/coffee"
)

type coffeeHandler struct {
	inventory *coffee.Inventory
}

func (ch coffeeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ch.get(w, r)
	case http.MethodPost:
		ch.post(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (ch coffeeHandler) get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	data, _ := json.Marshal(ch.inventory.GetAll())
	w.Write(data)
}

func (ch coffeeHandler) post(w http.ResponseWriter, r *http.Request) {
	var item coffee.Item
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse inventory item"))
		return
	}
	item = ch.inventory.Add(item)
	w.Header().Add("content-type", "application/json")

	data, _ := json.Marshal(item)
	w.Write(data)

}
