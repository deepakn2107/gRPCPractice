package main

import (
	"log"
	"net"

	pb "github.com/deepak2107/grpc-go-project/greet/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.GreetServiceServer
}

var add string = "0.0.0.0:5001"

func main() {
	lis, err := net.Listen("tcp", add)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listnenting on %s\n", add)

	s := grpc.NewServer()

	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
