package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/vinayreddy/experimental/calculatorpb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/grpc"
)

type server struct {}

func (s *server) Add(ctx context.Context, req *calculatorpb.AddRequest) (*calculatorpb.AddResponse, error) {
	fmt.Println("Got a new Add request")
	num1 := req.GetNum1()
	num2 := req.GetNum2()
	sum := num1 + num2
	result := &calculatorpb.AddResponse{Result: sum}
	log.Printf("Result: %v", protojson.Format(result))
	return result, nil
}

func main() {
	hostport := "0.0.0.0:8000"
	fmt.Println("Starting Calculator server on", hostport)
	lis, err := net.Listen("tcp", hostport)
	if err != nil {
		log.Fatalf("Error while listening : %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while serving : %v", err)
	}
}
