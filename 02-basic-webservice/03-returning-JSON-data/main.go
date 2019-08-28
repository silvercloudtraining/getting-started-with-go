package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		data := `{"name":"Tricia", "username":"tmacmillan"}`
		w.Write([]byte(data))
	})
	var err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
