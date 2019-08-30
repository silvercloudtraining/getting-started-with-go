package coffee

import "fmt"

type Item struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Quantity    float64 `json:"quantity"`
}

var nextID = 1

func NewInventory() *Inventory {
	return &Inventory{items: []Item{}}
}

type Inventory struct {
	items []Item
}

func (inv Inventory) Get(id int) (Item, error) {
	for _, item := range inv.items {
		if id == item.ID {
			return item, nil
		}
	}
	return Item{}, fmt.Errorf("Unable to find item with ID %v", id)
}

func (inv *Inventory) Add(item Item) Item {
	item.ID = nextID
	nextID++
	inv.items = append(inv.items, item)
	return item
}

func (inv Inventory) GetAll() []Item {
	return inv.items
}
