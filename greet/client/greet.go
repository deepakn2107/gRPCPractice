package main

import (
	"context"
	"log"

	pb "github.com/deepak2107/grpc-go-project/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreetWas invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Deepak",
	})
	if err != nil {
		log.Fatalf("could not greet: %v\n", err)
	}
	log.Printf("Greeting: %s\n", res.Result)
}
