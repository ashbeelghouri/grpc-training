package main

import (
	"fmt"
	"grpc/greet/greetpb"
	"log"

	"google.golang.org/grpc"
)


func main() {
	fmt.Println("Hello from client")
	client_connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer client_connection.Close()

	c := greetpb.NewGreetServiceClient(client_connection)

	fmt.Printf("Created Client: %v", c)

}