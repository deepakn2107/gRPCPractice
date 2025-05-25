package main

import (
	"context"
	"log"

	pb "github.com/deepak2107/grpc-go-project/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("update blog was invoked")
	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Steve Roggers",
		Title:    "A new title",
		Content:  "I Can do this all day",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error happpend while connecting to server for update: %v\n", err)
	}

	log.Println("Blogs was updated")
}
