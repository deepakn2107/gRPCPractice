package main

import (
	"io"
	"log"

	pb "github.com/deepak2107/grpc-go-project/calculator/proto"
)

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {

	log.Println("Average method at sever side invoked")

	count := 0
	var res int64 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Result: float64(res) / float64(count),
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving numbers: %d\n", req.Number)
		res += req.Number
		count++
	}

}
