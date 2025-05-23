package main

import (
	"context"
	"io"
	"log"

	pb "github.com/deepak2107/grpc-go-project/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Printf("doGreetManyTimes method got invoked")

	req := &pb.GreetRequest{
		FirstName: "Steve Roggers",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while sending request for GreetManytimes %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the stream %v\n", err)
		}

		log.Printf("Greet many times: %s\n", msg)
	}
}
