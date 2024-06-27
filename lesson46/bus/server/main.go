package main

import (
	"fmt"
	"log"
	"net"

	pb "weather_and_bus/bus/genproto/transport"
	"weather_and_bus/bus/service"
	"weather_and_bus/bus/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is listening on port 50051...")

	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	s := service.NewTransport(*postgres.NewTransportRepo(db))
	grpc := grpc.NewServer()
	pb.RegisterTransportServerServer(grpc, s)

	if err := grpc.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
