package main

import (
	"context"
	"fmt"
	"log"

	"github.com/vinayreddy/experimental/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	hostport := "0.0.0.0:8000"
	conn, err := grpc.Dial(hostport, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect to printer service: %v", err)
	}
	cc := calculatorpb.NewCalculatorClient(conn)
	cr := &calculatorpb.AddRequest{Num1: 5, Num2: 10}
	res, err := cc.Add(context.Background(), cr)
	if err != nil {
		log.Fatalf("Error calling calculator: %v", err)
	}
	fmt.Println(protojson.Format(res))
}
