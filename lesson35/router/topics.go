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
	NewTopicsRouter(tr, handler)

}
