package handler

import (
	"encoding/json"
	"fmt"
	"leetcode/model"
	"net/http"

	"github.com/gorilla/mux"
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
	vars := mux.Vars(r)

	if _, ok := vars["id"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error with getting id from query"))
		return
	}
	id := vars["id"]

	user, err := h.Users.GetUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error with getting data with the id: %s", err)))
		return
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error with converting data to json: %s", err)))
		return
	}
}

// update
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := &model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error with taking body from front: %s", err)))
		return
	}
	err = h.Users.UpdateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error with update data in database: %s", err)))
		return
	}
	w.Write([]byte("Updated user successfully"))
}

// delete
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if _, ok := vars["id"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error with getting id from query"))
		return
	}
	id := vars["id"]

	err := h.Users.DeleteUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error with deleting data from database: %s", err)))
		return
	}
	w.Write([]byte("Deleted user successfully"))
}
