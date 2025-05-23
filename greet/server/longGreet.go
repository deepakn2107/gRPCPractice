package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/deepak2107/grpc-go-project/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {

	log.Println("LongGreet function was invoked")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})

		}
		log.Printf("Receiving: %v\n", req)
		if err != nil {
			log.Fatalf("Error while reading the client stream: %v\n", err)
		}

		res += fmt.Sprintf("hello %s!\n", req.FirstName)
	}
}
