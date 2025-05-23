package main

import (
	"context"
	"log"

	pb "github.com/deepak2107/grpc-go-project/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient) {
	log.Println("doSqrt method got invoked from client side")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: 400})

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Printf("Probably sent wrong values: %s\n", e.Message())
			}
		} else {
			log.Fatalf("A not grpc error: %v\n", err)
		}
	}

	log.Printf("Sqrt Result: %f\n", res.Number)
}
