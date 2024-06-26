package main

import (
	"fmt"
	"log"

	pb "library/genprotos/generator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("tcp", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	c := pb.NewLibraryServiceClient(conn)
	fmt.Println(c)
}
