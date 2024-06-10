package router

import (
	"leetcode/handler"

	"github.com/gorilla/mux"
)

type ProblemsRouter struct {
	ProblemsRouter *mux.Router
	Handler        *handler.Handler
}

func NewProblemsRouter(pr *mux.Router, handler *handler.Handler) *ProblemsRouter {
	return &ProblemsRouter{ProblemsRouter: pr, Handler: handler}
}

func ProblemsServer(mainRouter *mux.Router, handler *handler.Handler) {
	p := mainRouter.PathPrefix("/Problems").Subrouter()
	r := NewProblemsRouter(p, handler)

	r.ProblemsRouter.HandleFunc("/Create", r.Handler.CreateProblem).Methods("POST")
	r.ProblemsRouter.HandleFunc("/Get", r.Handler.GetProblems).Methods("GET")
	r.ProblemsRouter.HandleFunc("/{id}", r.Handler.GetProblemByID).Methods("GET")
	r.ProblemsRouter.HandleFunc("/Update", r.Handler.UpdateProblem).Methods("PUT")
	r.ProblemsRouter.HandleFunc("/{id}", r.Handler.DeleteProblem).Methods("DELETE")
}
