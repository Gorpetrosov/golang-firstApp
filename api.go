package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type api struct {
	addr string
}

var users = []User{}

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//encode users slice to json
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (a *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {

	var payload User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	u := User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	}

	err = inserUser(u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//encode users slice to json
	error := json.NewEncoder(w).Encode(u)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func inserUser(u User) error {

	if u.FirstName == "" {
		return errors.New("First name is required")
	}

	if u.LastName == "" {
		return errors.New("Last name is required")
	}

	// storage validation
	for _, user := range users {
		if user.FirstName == u.FirstName && user.LastName == u.LastName {
			return errors.New("User already exists")
		}
	}

	users = append(users, u)
	return nil

}
