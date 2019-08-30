package user

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
}

var (
	users = []User{
		User{ID: 1, FirstName: "Arthur", LastName: "Dent", Username: "adent"},
		User{ID: 2, FirstName: "Tricia", LastName: "MacMillan", Username: "tmacmillan"},
		User{ID: 3, FirstName: "Zaphod", LastName: "Beeblebrox", Username: "zbeeblebrox"},
	}
	nextID = 4
)

func GetByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User not found with ID %v", id)
}

func GetAll() []User {
	return users
}

func Add(newUser User) User {
	newUser.ID = nextID
	nextID++
	users = append(users, newUser)
	return newUser
}

func Update(id int, u User) (User, error) {
	if id != u.ID {
		return User{}, errors.New("ID of updated user cannot change")
	}
	for i := range users {
		if users[i].ID == id {
			users[i] = u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User not found with ID %v", id)
}

func Delete(id int) error {
	for i := range users {
		if users[i].ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User not found with ID %v", id)
}
