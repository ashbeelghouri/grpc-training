package main

import (
	"context"
	"fmt"
	"grpc/calculator/calculationpb"
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

	c := calculationpb.NewCalculationServiceClient(client_connection)

	// fmt.Printf("Created Client: %v", c)

	doUnary(c)
	
}

func doUnary(c calculationpb.CalculationServiceClient){
	req := &calculationpb.CalculationRequest{
		Calculation: &calculationpb.Calculation{
			A: 3,
			B: 10,
		},
	}

	res, err := c.Calculate(context.Background(), req)

	if err != nil {
		log.Printf("error while calling Greet rpc %v", err)
	}

	log.Printf("Response from Greet %v", res.Result)
}