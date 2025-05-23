package main

import (
	"context"
	"log"

	pb "github.com/deepak2107/grpc-go-project/calculator/proto"
)

func doAverage(c pb.CalculatorServiceClient) {

	reqs := []*pb.AverageRequest{
		{Number: 2},
		{Number: 2},
		{Number: 2},
		{Number: 2},
		{Number: 2},
		{Number: 2},
	}

	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Average Request, %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response%v\n", err)
	}

	log.Printf("Average: %f", res.Result)
}
