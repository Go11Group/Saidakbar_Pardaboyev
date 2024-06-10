package handler

import (
	"encoding/json"
	"fmt"
	"leetcode/model"
	"net/http"
)

// create
func (h *Handler) CreateUsers(w http.ResponseWriter, r *http.Request) {
	newUser := &model.User{}
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error with Getting user from frontend: %s", err)))
		return
	}

	err = h.Users.CreateUser(newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error with Creating new user in database: %s", err)))
		return
	}
	w.Write([]byte("Successfully Created New User"))
}

// readAll
func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	filter := &model.FilterUser{}
	if params.Has("fullname") {
		fullname := params.Get("fullname")
		filter.Fullname = &fullname
	}
	if params.Has("username") {
		username := params.Get("username")
		filter.Username = &username
	}

	users, err := h.Users.GetUsers(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error getting data from database: %s", err)))
		return
	}

	err = json.NewEncoder(w).Encode(&users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error converting data to json: %s", err)))
		return
	}
}

// read By One
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {

}

// update
func (h *Handler) UpdateUsers(w http.ResponseWriter, r *http.Request) {

}

// delete
func (h *Handler) DeleteUsers(w http.ResponseWriter, r *http.Request) {

}
