package main

import (
	"context"
	"log"

	pb "github.com/deepak2107/grpc-go-project/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("deleteBlog method invoked")

	req := &pb.BlogId{Id: id}

	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Printf("Error happend while reading: %v\n", err)
	}

	log.Printf("Blog was deleted: %v\n", res)

}
