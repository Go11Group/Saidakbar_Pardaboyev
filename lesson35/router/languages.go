package router

import (
	"leetcode/handler"

	"github.com/gorilla/mux"
)

type LanguagesRouter struct {
	LanguagesRouter *mux.Router
	Handler         *handler.Handler
}

func NewLanguageRouter(l *mux.Router, handler *handler.Handler) *LanguagesRouter {
	return &LanguagesRouter{LanguagesRouter: l, Handler: handler}
}

func LanguageServer(mainRouter *mux.Router, handler *handler.Handler) {
	l := mainRouter.PathPrefix("/Language").Subrouter()
	NewLanguageRouter(l, handler)

}
