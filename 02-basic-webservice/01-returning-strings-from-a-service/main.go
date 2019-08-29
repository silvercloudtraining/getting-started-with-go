package main

import (
	"net/http"
)

const port = ":3000"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Gophers!"))
	})
	http.ListenAndServe(port, nil)
}
