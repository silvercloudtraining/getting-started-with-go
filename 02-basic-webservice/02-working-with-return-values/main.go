package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Gophers"))
	})
	var err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
