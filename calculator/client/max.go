package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/deepak2107/grpc-go-project/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax method got invoked at client side")

	waitc := make(chan struct{})

	reqs := []*pb.MaxRequest{
		{Number: 1},
		{Number: 2},
		{Number: 30},
		{Number: 4},
		{Number: 1},
	}
	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream to server")
	}
	go func() {
		for _, req := range reqs {

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
				log.Fatalf("Error while receiving response: %v\n", err)
			}

			log.Printf("Received max number: %d\n", res.Result)

		}
		close(waitc)
	}()

	<-waitc
}
