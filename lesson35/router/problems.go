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
	NewProblemsRouter(p, handler)

}
