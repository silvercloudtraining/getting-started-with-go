package routing

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/silvercloudtraining/coffeeservice/user"
)

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
		}
	} else {
		id, _ := strconv.Atoi(matches[1])
		getOne(id, w, r)
	}
}

func getOne(id int, w http.ResponseWriter, r *http.Request) {
	u, err := user.GetByID(id)
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
	users := user.GetAll()
	w.Header().Add("content-type", "application/json")
	enc := json.NewEncoder(w)
	err := enc.Encode(users)
	if err != nil {
		log.Print(err)
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	var newUser user.User
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to parse incoming user data"))
		return
	}
	newUser = user.Add(newUser)
	data, err := json.Marshal(newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.Write(data)
}
