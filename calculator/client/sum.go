package main

import (
	"context"
	"log"

	pb "github.com/deepak2107/grpc-go-project/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Printf("doSum method invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNum:  10,
		SecondNum: 11,
	})
	if err != nil {
		log.Printf("could not send sum request %v\n", err)
	}

	log.Printf("Result after calling the sever for sum: %d\n", res.Result)
}
