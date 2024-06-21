package main

import (
	"client/api"
	"client/api/handler"
)

func main() {

	requestForUsers()
}

func requestForUsers() {

	h := handler.NewHandler()
	server := api.CreateServer(h)
	server.ListenAndServe()
}
