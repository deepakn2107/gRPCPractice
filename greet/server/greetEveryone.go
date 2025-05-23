package main

import (
	"io"
	"log"

	pb "github.com/deepak2107/grpc-go-project/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {

	log.Println("GreetEveryone method got invoked")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading request from client %v\n", err)
		}

		res := "Hello" + " " + req.FirstName

		err = stream.Send(&pb.GreetResponse{Result: res})
		if err != nil {
			log.Fatalf("Error while sending response to client %v\n", err)
		}
	}
}
