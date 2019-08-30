package routing

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/silvercloudtraining/coffeeservice/coffee"
)

type coffeeInventory interface {
	Add(newItem coffee.Item) coffee.Item
	GetAll() []coffee.Item
	Get(id int) (coffee.Item, error)
}

func NewCoffeeHandler(inventory coffeeInventory) *coffeeHandler {
	return &coffeeHandler{
		inventory:   inventory,
		itemPattern: regexp.MustCompile(`^/coffee/(\d+)$`),
	}
}

type coffeeHandler struct {
	inventory   coffeeInventory
	itemPattern *regexp.Regexp
}

func (ch coffeeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	matches := ch.itemPattern.FindStringSubmatch(r.URL.Path)
	if len(matches) == 0 {
		switch r.Method {
		case http.MethodGet:
			ch.get(w, r)
		case http.MethodPost:
			ch.post(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	} else {
		id, _ := strconv.Atoi(matches[1])
		switch r.Method {
		case http.MethodGet:
			ch.getOne(id, w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func (ch coffeeHandler) getOne(id int, w http.ResponseWriter, r *http.Request) {
	item, err := ch.inventory.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data, _ := json.Marshal(item)
	w.Header().Add("content-type", "application/json")
	w.Write(data)
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
