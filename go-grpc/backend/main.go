package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/grpc/api"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreetServiceServer
}

func (s server) Hello(ctx context.Context, hr *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Println("request to grpc!!", hr.GetMessage())
	return &pb.HelloResponse{Message: hr.GetMessage()}, nil
}

func (s server) User(ctx context.Context, ur *pb.UserRequest) (*pb.UserResponse, error) {
	log.Println("request for a user from client")
	user := pb.User{Id: "1234", Name: "John", Age: 32}
	return &pb.UserResponse{User: &user}, nil
}

func main() {
	const PORT = 8080
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", PORT))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listening on localhost:%d", PORT)
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, server{})
	grpcServer.Serve(lis)
}
