package main

import (
	"context"
	"log"

	pb "github.com/deepak2107/grpc-go-project/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("CreateBlog method invoked")

	blog := &pb.Blog{
		AuthorId: "Deepak",
		Title:    "My first blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Error while connecting to server%v\n", err)
	}

	log.Printf("Blog has been created: %v", res.Id)

	return res.Id
}
