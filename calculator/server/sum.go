package main

import (
	"context"
	"log"

	pb "github.com/deepak2107/grpc-go-project/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function is called with params %d, %d", in.FirstNum, in.SecondNum)
	return &pb.SumResponse{
		Result: in.FirstNum + in.SecondNum,
	}, nil
}
