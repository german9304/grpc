package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/grpc/api"
	"google.golang.org/grpc"
)

func main() {
	const PORT = 8080
	// calls server
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", PORT), grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewGreetServiceClient(conn)

	response, err := client.Hello(context.Background(), &pb.HelloRequest{Message: "hello World"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("server response %s \n", response)

	userResponse, err := client.User(context.Background(), &pb.UserRequest{})
	if err != nil {
		log.Println(err)
	}
	user := userResponse.GetUser()
	log.Printf("age: %d name: %s id: %s \n", user.GetAge(), user.GetName(), user.GetId())
}
