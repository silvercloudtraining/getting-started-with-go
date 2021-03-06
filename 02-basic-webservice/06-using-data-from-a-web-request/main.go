package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const port = ":3000"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		name := r.URL.Query().Get("name")
		username := r.URL.Query().Get("username")
		m := map[string]string{"name": name,
			"username": username,
		}
		enc := json.NewEncoder(w)
		err := enc.Encode(m)
		if err != nil {
			log.Print(err)
		}

	})
	var err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
