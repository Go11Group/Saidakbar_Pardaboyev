package main

import (
	"context"
	pb "library/genprotos/generator"
	"log"
	"net"

	"google.golang.org/grpc"
)

type LibraryService struct {
	pb.UnimplementedLibraryServiceServer
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(listener)
	}

	s := LibraryService{}
	grpc := grpc.NewServer()
	pb.RegisterLibraryServiceServer(grpc, &s)

	if err := grpc.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

func (l *LibraryService) AddBook(ctx context.Context, req *pb.AddBookRequest) (*pb.AddBookResponse, error) {

}

func (l *LibraryService) SearchBook(ctx context.Context, req *pb.SearchBookRequest) (*pb.SearchBookResponse, error) {

}

func (l *LibraryService) BorrowBook(ctx context.Context, req *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {

}
