package router

import (
	"leetcode/handler"

	"github.com/gorilla/mux"
)

type TopicsRouter struct {
	TopicsRouter *mux.Router
	Handler      *handler.Handler
}

func NewTopicsRouter(tr *mux.Router, handler *handler.Handler) *TopicsRouter {
	return &TopicsRouter{TopicsRouter: tr, Handler: handler}
}

func TopicsServer(mainRouter *mux.Router, handler *handler.Handler) {
	tr := mainRouter.PathPrefix("/Topics").Subrouter()
	r := NewTopicsRouter(tr, handler)

	r.TopicsRouter.HandleFunc("/Create", r.Handler.CreateTopic).Methods("POST")
	r.TopicsRouter.HandleFunc("/Get", r.Handler.GetTopics).Methods("GET")
	r.TopicsRouter.HandleFunc("/{id}", r.Handler.GetTopicByID).Methods("GET")
	r.TopicsRouter.HandleFunc("/Update", r.Handler.UpdateTopic).Methods("PUT")
	r.TopicsRouter.HandleFunc("/{id}", r.Handler.DeleteTopic).Methods("DELETE")
}
