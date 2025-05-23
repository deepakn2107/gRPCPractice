package main

import (
	"context"
	"log"
	"time"

	pb "github.com/deepak2107/grpc-go-project/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet with Deadline was inovked with %v\n", in)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("the client calceled the request!")
			return nil, status.Error(codes.Canceled, "The client cancelled the request")
		}
		time.Sleep(1 * time.Second)

	}

	return &pb.GreetResponse{
		Result: "hello " + in.FirstName,
	}, nil
}
