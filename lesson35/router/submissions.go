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
	NewSubmissionsRouter(sb, handler)

}
