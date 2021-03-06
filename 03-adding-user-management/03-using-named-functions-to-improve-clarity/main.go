package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const port = ":3000"

type user struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
}

var (
	users = []user{
		user{ID: 1, FirstName: "Arthur", LastName: "Dent", Username: "adent"},
		user{ID: 2, FirstName: "Tricia", LastName: "MacMillan", Username: "tmacmillan"},
		user{ID: 3, FirstName: "Zaphod", LastName: "Beeblebrox", Username: "zbeeblebrox"},
	}
	nextID = 4
)

func main() {
	http.HandleFunc("/users", userHandler)
	var err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		get(w, r)
	case http.MethodPost:
		post(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	users := getUsers()
	w.Header().Add("content-type", "application/json")
	enc := json.NewEncoder(w)
	err := enc.Encode(users)
	if err != nil {
		log.Print(err)
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	var newUser user
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to parse incoming user data"))
		return
	}
	newUser = addUser(newUser)
	data, err := json.Marshal(newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.Write(data)
}

func getUsers() []user {
	return users
}

func addUser(newUser user) user {
	newUser.ID = nextID
	nextID++
	users = append(users, newUser)
	return newUser
}
