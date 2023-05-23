package rpc

import (
	"fmt"
	"github.com/alima12/Blog-Go/service/compiles"
	"github.com/alima12/Blog-Go/service/manager"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func StartRpcServer() {
	_ = godotenv.Load(".env")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("RPC_PORT")))
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()

	servicePost := &manager.PostService{}
	compiles.RegisterPostServiceServer(server, servicePost)
	serviceAuth := &manager.AuthenticationService{}
	compiles.RegisterAuthenticationServer(server, serviceAuth)
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("imposible to serve: %s", err)
	}

}
