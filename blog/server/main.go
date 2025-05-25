package main

import (
	"context"
	"log"
	"net"

	pb "github.com/deepak2107/grpc-go-project/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.BlogServiceServer
}

var collection *mongo.Collection
var add string = "localhost:5001"

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatalf("Error while creating mongo db connection %v\n", err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatalf("Error while creating mongo db connection %v\n", err)
	}

	collection = client.Database("blogdb").Collection("blog")
	lis, err := net.Listen("tcp", add)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", add)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})
	reflection.Register(s)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
