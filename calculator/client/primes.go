package main

import (
	"context"
	"io"
	"log"

	pb "github.com/deepak2107/grpc-go-project/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Printf("doPrimes method invoked from client side")

	req := &pb.PrimeRequest{
		Number: 120,
	}

	stream, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while sending stream request for Primes %v\n", err)
	}

	for {
		number, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream %v\n", err)
		}

		log.Printf("Got the prime number: %d\n", number.Number)

	}
}
