package rpc

import (
	"context"
	"errors"
	"github.com/alima12/Blog-Go/database"
	"github.com/alima12/Blog-Go/models"
	compiles "github.com/alima12/Blog-Go/service/compiles/proto"
	compiles2 "github.com/alima12/Blog-Go/service/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
)

type PostService struct {
	compiles.UnimplementedPostServiceServer
}

func (ps *PostService) GetSinglePost(context.Context, *compiles2.RetrievePost) (*compiles2.SinglePostResponse, error) {
	db := database.GetDB()
	//id := c.Param("id")
	var post models.Post
	result := db.First(&post, 1)
	if result.Error != nil {
		return nil, errors.New("post not found")
	}
	response := &compiles2.SinglePostResponse{
		Title:  post.Title,
		Body:   post.Body,
		UserID: int32(post.UserID),
	}
	return response, nil
}

func StartRpcServer() {
	_ = godotenv.Load(".env")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()

	service := &PostService{}
	compiles.RegisterPostServiceServer(server, service)
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("imposible to serve: %s", err)
	}

}
