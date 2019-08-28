package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	myInt := 42
	myFloat := 3.14
	myBool := true
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		type response struct {
			MyInt   int
			MyFloat float64
			MyBool  bool
		}
		res := response{
			MyInt:   myInt,
			MyFloat: myFloat,
			MyBool:  myBool,
		}
		enc := json.NewEncoder(w)
		err := enc.Encode(res)
		if err != nil {
			log.Print(err)
		}

	})
	var err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
