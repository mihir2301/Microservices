package main

import (
	"common/api"
	"context"
	"log"

	"google.golang.org/grpc"
)

type grpcHandler struct {
	api.UnimplementedOrderServiceServer
}

func NewGRPCHandler(grpcServer *grpc.Server) {
	handler := &grpcHandler{}
	api.RegisterOrderServiceServer(grpcServer, handler)
}
func (h *grpcHandler) Createorder(context.Context, *api.CreateOrderRequest) (*api.Order, error) {
	log.Println("New order recieved")
	o := &api.Order{
		ID: "42",
	}
	return o, nil
}
