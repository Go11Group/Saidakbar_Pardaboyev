package main

import (
	"fmt"
	"log"
	"net"
	pbT "weather_and_bus/genproto/transport"
	pbW "weather_and_bus/genproto/weather"
	service "weather_and_bus/services"
	"weather_and_bus/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	t := service.NewTransport(postgres.NewTransportRepo(db))
	w := service.NewWeather(postgres.NewWeatherRepo(db))

	grpcServer := grpc.NewServer()
	pbW.RegisterWeatherServiceServer(grpcServer, w)
	pbT.RegisterTransportServerServer(grpcServer, t)

	fmt.Println("Server is listening on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
