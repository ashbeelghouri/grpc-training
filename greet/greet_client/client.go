package main

import (
	"context"
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

	// fmt.Printf("Created Client: %v", c)

	doUnary(c)
	
}

func doUnary(c greetpb.GreetServiceClient){
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			Firstname: "Ashbeel",
			Lastname: "Ghouri",
		},
	}

	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Printf("error while calling Greet rpc %v", err)
	}

	log.Printf("Response from Greet %v", res.Result)
}