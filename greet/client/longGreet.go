package main

import (
	"context"
	"log"
	"time"

	pb "github.com/deepak2107/grpc-go-project/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Captain America"},
		{FirstName: "Iron Man"},
		{FirstName: "Blackwidow"},
		{FirstName: "Thor"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling long Greet method %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response%v\n", err)
	}

	log.Printf("Long Greet: %s\n", res.Result)
}
