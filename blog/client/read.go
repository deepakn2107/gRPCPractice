package main

import (
	"context"
	"log"

	pb "github.com/deepak2107/grpc-go-project/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("readBlog method invoked")

	req := &pb.BlogId{Id: id}

	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Printf("Error happend while reading: %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)

	return res
}
