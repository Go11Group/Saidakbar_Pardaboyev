package router

import (
	"leetcode/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateNewServer(handler *handler.Handler) *http.Server {
	LeetcodeRouter := mux.NewRouter().PathPrefix("/LeetCode").Subrouter()

	UsersServer(LeetcodeRouter, handler)
	ProblemsServer(LeetcodeRouter, handler)
	ProblemsTopicsServer(LeetcodeRouter, handler)
	LanguageServer(LeetcodeRouter, handler)
	TopicsServer(LeetcodeRouter, handler)
	SubmissionsServer(LeetcodeRouter, handler)

	return &http.Server{
		Addr:    ":8080",
		Handler: LeetcodeRouter,
	}
}
