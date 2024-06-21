package handler

import "net/http"

type Handler struct {
	Client *http.Client
}

func NewHandler() *Handler {
	client := http.Client{}
	return &Handler{Client: &client}
}
