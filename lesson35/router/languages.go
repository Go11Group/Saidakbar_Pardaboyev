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
	r := NewLanguageRouter(l, handler)

	r.LanguagesRouter.HandleFunc("/Create", r.Handler.CreateLanguages).Methods("POST")
	r.LanguagesRouter.HandleFunc("/Get", r.Handler.GetLanguages).Methods("GET")
	r.LanguagesRouter.HandleFunc("/{id}", r.Handler.GetLanguageByID).Methods("GET")
	r.LanguagesRouter.HandleFunc("/Update", r.Handler.UpdateLanguage).Methods("PUT")
	r.LanguagesRouter.HandleFunc("/{id}", r.Handler.DeleteLanguage).Methods("DELETE")
}
