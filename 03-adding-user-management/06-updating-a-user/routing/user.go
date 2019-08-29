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
			getUsers(w, r)
		case http.MethodPost:
			postUsers(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	} else {
		id, _ := strconv.Atoi(matches[1])
		switch r.Method {
		case http.MethodGet:
			getOneUser(id, w, r)
		case http.MethodPut:
			updateUser(id, w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func updateUser(id int, w http.ResponseWriter, r *http.Request) {
	var u user.User
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	u, err = user.Update(id, u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	data, _ := json.Marshal(u)
	w.Header().Add("content-type", "application/json")
	w.Write(data)
}

func getOneUser(id int, w http.ResponseWriter, r *http.Request) {
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

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := user.GetAll()
	w.Header().Add("content-type", "application/json")
	enc := json.NewEncoder(w)
	err := enc.Encode(users)
	if err != nil {
		log.Print(err)
	}
}

func postUsers(w http.ResponseWriter, r *http.Request) {
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
