package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

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
	http.HandleFunc("/users/", userHandler)
	var err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	regex := regexp.MustCompile(`^/users/(\d+)$`)
	matches := regex.FindStringSubmatch(r.URL.Path)
	if len(matches) == 0 {
		switch r.Method {
		case http.MethodGet:
			get(w, r)
		case http.MethodPost:
			post(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	} else {
		id, _ := strconv.Atoi(matches[1])
		getOne(id, w, r)
	}
}

func getOne(id int, w http.ResponseWriter, r *http.Request) {
	u, err := getByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("content-type", "application/json")
	enc := json.NewEncoder(w)
	err = enc.Encode(u)
	if err != nil {
		log.Print(err)
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
	w.Write(data)
	return
}

func getByID(id int) (user, error) {
	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}
	return user{}, fmt.Errorf("User not found with ID %v", id)
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
