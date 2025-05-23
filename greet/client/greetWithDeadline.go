package main

import (
	"context"
	"log"
	"time"

	pb "github.com/deepak2107/grpc-go-project/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	req := &pb.GreetRequest{
		FirstName: " Natasha Romanoff",
	}

	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Printf("Deadline exceeded")
				return
			} else {
				log.Fatalf("Unexpected error %v\n", err)
			}
		}
	}

	log.Printf("GreetWithDeadline: %s\n", res.Result)
}
