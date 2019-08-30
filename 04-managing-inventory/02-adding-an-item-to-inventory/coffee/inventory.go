package coffee

type Item struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Quantity    float64 `json:"quantity"`
}

var nextID = 1

type Inventory struct {
	items []Item
}

func (inv *Inventory) Add(item Item) Item {
	item.ID = nextID
	nextID++
	if inv.items == nil {
		inv.items = []Item{item}
	} else {
		inv.items = append(inv.items, item)
	}
	return item
}

func (inv Inventory) GetAll() []Item {
	return inv.items
}
