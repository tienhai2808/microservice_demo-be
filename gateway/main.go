package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/tienhai2808/microservice_demo-be/common"
	"github.com/tienhai2808/microservice_demo-be/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	orderServiceAddr = "localhost:8081"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or failed to load it")
	}
	httpAddr := common.EnvString("HTTP_ADDR", ":3000")

	conn, err := grpc.Dial(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()
	log.Println("Dialing orders service at ", orderServiceAddr)

	client := api.NewOrderServiceClient(conn)

	mux := http.NewServeMux()
	handler := NewHttpHandler(client)
	handler.registerRoutes(mux)

	log.Printf("Starting HTTP server at %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server")
	}
}
