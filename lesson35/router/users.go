package router

import (
	"leetcode/handler"

	"github.com/gorilla/mux"
)

type UsersRouter struct {
	UsersRouter *mux.Router
	Handler     *handler.Handler
}

func NewUsersRouter(ur *mux.Router, handler *handler.Handler) *UsersRouter {
	return &UsersRouter{UsersRouter: ur, Handler: handler}
}

func UsersServer(LeetcodeRouter *mux.Router, handler *handler.Handler) {
	ur := LeetcodeRouter.PathPrefix("/Users").Subrouter()
	r := NewUsersRouter(ur, handler)

	r.UsersRouter.HandleFunc("/Create", r.Handler.CreateUsers).Methods("POST")
	r.UsersRouter.HandleFunc("/Get", r.Handler.GetUsers).Methods("GET")
	// r.UsersRouter.HandleFunc("/{id}", r.Handler.GetUsers).Methods("PUT")
}
