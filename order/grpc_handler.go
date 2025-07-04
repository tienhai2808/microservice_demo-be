package main

import (
	"context"
	"log"

	"github.com/tienhai2808/microservice_demo-be/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	api.UnimplementedOrderServiceServer
	service OrderService
}

func NewGRPCHandler(grpcServer *grpc.Server, service OrderService) {
	handler := &grpcHandler{
		service: service,
	}
	api.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, payload *api.CreateOrderRequest) (*api.Order, error) {
	log.Printf("New order received!!!, Order: %v", payload)
	order := &api.Order{
		ID: "42",
	}
	return order, nil
}
