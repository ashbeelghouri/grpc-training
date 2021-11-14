package main

import (
	"context"
	"fmt"
	"grpc/greet/greetpb"
	"io"
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
	doServerStreaming(c)

}

func doServerStreaming(c greetpb.GreetServiceClient) {
	log.Println("GreetManyTimes Streaming Method")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			Firstname: "Ashbeel",
			Lastname: "Ghouri",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet Many rpc %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break;
		}

		if err != nil {
			log.Fatalf("error while recieving many greets %v", err)
		}

		log.Printf("Response from GreetManyTimes %v", msg.GetResult())
	}
}

func doUnary(c greetpb.GreetServiceClient){
	log.Println("Greet Unary Method")
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