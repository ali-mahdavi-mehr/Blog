package rpc

import (
	"github.com/alima12/Blog-Go/service/compiles"
	"github.com/alima12/Blog-Go/service/service_models"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
)

func StartRpcServer() {
	_ = godotenv.Load(".env")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()

	service := &service_models.PostService{}
	compiles.RegisterPostServiceServer(server, service)
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("imposible to serve: %s", err)
	}

}
