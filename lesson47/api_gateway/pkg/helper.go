package pkg

import (
	pbT "api_gateway/genproto/transport"
	pbW "api_gateway/genproto/weather"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateTransportServerClient() *pbT.TransportServerClient {
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	c := pbT.NewTransportServerClient(conn)
	return &c
}

func CreateWeatherServerClient() *pbW.WeatherServiceClient {
	conn, err := grpc.NewClient(":50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	c := pbW.NewWeatherServiceClient(conn)
	return &c
}
