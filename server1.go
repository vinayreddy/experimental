package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/vinayreddy/experimental/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

type server struct {
	printerClient calculatorpb.PrinterClient
}

func (s *server) Add(ctx context.Context, req *calculatorpb.AddRequest) (*calculatorpb.AddResponse, error) {
	num1 := req.GetNum1()
	num2 := req.GetNum2()
	sum := num1 + num2
	result := &calculatorpb.AddResponse{Result: sum}
	pr := &calculatorpb.PrintRequest{Num: sum}
	if _, err := s.printerClient.Print(context.Background(), pr); err != nil {
		log.Fatalf("Error calling printer: %v", err)
	}
	fmt.Println(protojson.Format(result))
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
	pa := "0.0.0.0:8080"
	conn, err := grpc.Dial(pa, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect to printer service: %v", err)
	}
	pc := calculatorpb.NewPrinterClient(conn)
	calculatorpb.RegisterCalculatorServer(s, &server{printerClient: pc})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while serving : %v", err)
	}
}
