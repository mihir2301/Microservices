package main

import (
	"common"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {

	grpcServer := grpc.NewServer()

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Println("Error in connecting to grpc", err)
		return
	}

	store := NewStore()
	svc := NewService(store)

	svc.CreateOrder(context.Background())

	err = grpcServer.Serve(l)
	if err != nil {
		fmt.Println("error in serving grpc")
		return
	}
}
