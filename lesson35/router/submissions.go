package router

import (
	"leetcode/handler"

	"github.com/gorilla/mux"
)

type SubmissionsRouter struct {
	SubmissionsRouter *mux.Router
	Handler           *handler.Handler
}

func NewSubmissionsRouter(sb *mux.Router, handler *handler.Handler) *SubmissionsRouter {
	return &SubmissionsRouter{SubmissionsRouter: sb, Handler: handler}
}

func SubmissionsServer(mainRouter *mux.Router, handler *handler.Handler) {
	sb := mainRouter.PathPrefix("/Submissions").Subrouter()
	r := NewSubmissionsRouter(sb, handler)

	r.SubmissionsRouter.HandleFunc("/Create", r.Handler.CreateSubmission).Methods("POST")
	r.SubmissionsRouter.HandleFunc("/Get", r.Handler.GetSubmissions).Methods("GET")
	r.SubmissionsRouter.HandleFunc("/{id}", r.Handler.GetSubmissionByID).Methods("GET")
	r.SubmissionsRouter.HandleFunc("/Update", r.Handler.UpdateSubmission).Methods("PUT")
	r.SubmissionsRouter.HandleFunc("/{id}", r.Handler.DeleteSubmission).Methods("DELETE")
}
