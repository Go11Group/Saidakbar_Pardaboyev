package router

import (
	"leetcode/handler"

	"github.com/gorilla/mux"
)

type ProblemsTopicsRouter struct {
	ProblemsTopicsRouter *mux.Router
	Handler              *handler.Handler
}

func NewProblemsTopicsRouter(pt *mux.Router, handler *handler.Handler) *ProblemsTopicsRouter {
	return &ProblemsTopicsRouter{ProblemsTopicsRouter: pt, Handler: handler}
}

func ProblemsTopicsServer(mainRouter *mux.Router, handler *handler.Handler) {
	pt := mainRouter.PathPrefix("/ProblemsTopics").Subrouter()
	r := NewProblemsTopicsRouter(pt, handler)

	r.ProblemsTopicsRouter.HandleFunc("/Create", r.Handler.CreateTopicProblem).Methods("POST")
	r.ProblemsTopicsRouter.HandleFunc("/Get", r.Handler.GetProblemsTopics).Methods("GET")
	r.ProblemsTopicsRouter.HandleFunc("/{id}", r.Handler.GetProblemTopicsById).Methods("GET")
	r.ProblemsTopicsRouter.HandleFunc("/Update", r.Handler.UpdateProblemTopic).Methods("PUT")
	r.ProblemsTopicsRouter.HandleFunc("/{id}", r.Handler.DeleteProblemTopic).Methods("DELETE")
}
