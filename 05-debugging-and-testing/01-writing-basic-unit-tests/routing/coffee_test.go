package routing

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/silvercloudtraining/coffeeservice/coffee"
)

type mockInventory struct {
	callCount map[string]int
}

func newMockInventory() mockInventory {
	return mockInventory{callCount: map[string]int{}}
}

func (mi mockInventory) Add(item coffee.Item) coffee.Item {
	return coffee.Item{}
}
func (mi mockInventory) GetAll() []coffee.Item {
	return nil
}
func (mi mockInventory) Get(id int) (coffee.Item, error) {
	mi.callCount["Get"] = mi.callCount["Get"] + 1
	return coffee.Item{}, nil
}

func TestCoffeeHandler(t *testing.T) {
	inv := newMockInventory()
	ch := NewCoffeeHandler(inv)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "http://localhost:3000/coffee", nil)
	t.Run("get", func(t *testing.T) {
		t.Run("calls correct method on inventory", func(t *testing.T) {
			ch.getOne(1, w, r)
			if inv.callCount["Get"] != 1 {
				t.Error("inventory.Get not called expected number of times")
			}
		})
	})

}
