package main

import (
	"io"
	"log"

	pb "github.com/deepak2107/grpc-go-project/calculator/proto"
)

func (*Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max was invoked")
	var maximum int64 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
		log.Printf("Receiving numbers: %d\n", req.Number)
		if number := req.Number; number > maximum {
			maximum = number

		}
		err = stream.Send(&pb.MaxResponse{
			Result: maximum,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
}
