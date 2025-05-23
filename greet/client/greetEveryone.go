package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/deepak2107/grpc-go-project/greet/proto"
)

func doGreetEveryOne(c pb.GreetServiceClient) {
	log.Printf("invoked doGreetEveryone method from client side")

	reqs := []*pb.GreetRequest{
		{FirstName: "Captain America"},
		{FirstName: "Iron Man"},
		{FirstName: "Blackwidow"},
		{FirstName: "Thor"},
	}

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("sending request: %v\n", req)

			err := stream.Send(req)
			if err != nil {
				log.Fatalf("Error while creating stream: %v\n", err)
			}

			time.Sleep(1 * time.Second)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error while receiving : %v\n", err)
			}

			log.Printf("Received: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
