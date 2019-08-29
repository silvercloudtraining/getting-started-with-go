package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type user struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
}

var users = [3]user{
	user{ID: 1, FirstName: "Arthur", LastName: "Dent", Username: "adent"},
	user{ID: 2, FirstName: "Tricia", LastName: "MacMillan", Username: "tmacmillan"},
	user{ID: 3, FirstName: "Zaphod", LastName: "Beeblebrox", Username: "zbeeblebrox"},
}

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		enc := json.NewEncoder(w)
		err := enc.Encode(users)
		if err != nil {
			log.Print(err)
		}

	})
	var err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
