package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/vinayreddy/experimental/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Print(context context.Context, req *calculatorpb.PrintRequest) (*calculatorpb.PrintResponse, error) {
	fmt.Printf("Print request: %v\n", req.GetNum())
	result := &calculatorpb.PrintResponse{}
	return result, nil
}

func main() {
	hostport := "0.0.0.0:8001"
	fmt.Println("Starting Printer server on", hostport)
	lis, err := net.Listen("tcp", hostport)

	if err != nil {
		log.Fatalf("Error while listening : %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterPrinterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while serving : %v", err)
	}
}
