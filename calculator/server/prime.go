package main

import (
	"log"

	pb "github.com/deepak2107/grpc-go-project/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Prime Method got invoked in server side, %v\n", in)

	number := in.Number

	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{
				Number: divisor,
			})
			number /= divisor
		} else {
			divisor++
		}

	}

	return nil
}
