package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Gophers!"))
	})
	http.ListenAndServe(":3000", nil)
}
