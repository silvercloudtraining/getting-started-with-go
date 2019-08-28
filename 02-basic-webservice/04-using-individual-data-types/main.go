package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	myInt := 42
	myFloat := 3.14
	myBool := true
	myComplex := complex(3, 4)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		format := `{
			"int": %v,
			"float": %v,
			"bool": %v,
			"complex": "%v"
		}`
		fmt.Fprintf(w, format, myInt, myFloat, myBool, myComplex)
	})
	var err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
