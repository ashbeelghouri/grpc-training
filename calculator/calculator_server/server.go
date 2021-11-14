package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"grpc/calculator/calculationpb"

	"google.golang.org/grpc"
)

type server struct{
	calculationpb.UnimplementedCalculationServiceServer
}

func (*server) Calculate(ctx context.Context, req *calculationpb.CalculationRequest) (*calculationpb.CalculationResponse, error) {

	fmt.Printf("Greet function was invoked with %v\n", req)
	a := req.GetCalculation().GetA()
	b := req.GetCalculation().GetB()
	result := a + b
	res := &calculationpb.CalculationResponse{
		Result: result,
	}
	return res, nil

}

func main() {
	fmt.Println("Hello world");
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()

	calculationpb.RegisterCalculationServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	
}