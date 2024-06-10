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
	NewProblemsTopicsRouter(pt, handler)

}
