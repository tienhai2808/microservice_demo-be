package main

import (
	"context"
	"log"
	"net"

	"github.com/tienhai2808/microservice_demo-be/common"
	"google.golang.org/grpc"
)

func main() {
	grpcAddr := common.EnvString("GRPC_ADDR", "localhost:8081")

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer l.Close()

	grpcServer := grpc.NewServer()
	store := NewOrderStore()
	service := NewOrderService(store)
	NewGRPCHandler(grpcServer, service)

	service.CreateOrder(context.Background())

	log.Println("GRPC started at: ", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}
