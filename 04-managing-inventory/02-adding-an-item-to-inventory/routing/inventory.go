package routing

import (
	"encoding/json"
	"net/http"

	"github.com/silvercloudtraining/coffeeservice/inventory"
)

type inventoryHandler struct {
	inventory *inventory.Inventory
}

func (ih inventoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ih.get(w, r)
	case http.MethodPost:
		ih.post(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (ih inventoryHandler) get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	data, _ := json.Marshal(ih.inventory.GetAll())
	w.Write(data)
}

func (ih inventoryHandler) post(w http.ResponseWriter, r *http.Request) {
	var item inventory.Item
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse inventory item"))
		return
	}
	item = ih.inventory.Add(item)
	w.Header().Add("content-type", "application/json")

	data, _ := json.Marshal(item)
	w.Write(data)

}
