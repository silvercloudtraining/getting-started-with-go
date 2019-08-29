package main

import (
	"log"
	"net/http"
)

const port = ":3000"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		data := `{"name":"Tricia", "username":"tmacmillan"}`
		w.Write([]byte(data))
	})
	var err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
