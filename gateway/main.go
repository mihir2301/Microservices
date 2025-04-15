package main

import (
	"common"
	"common/api"
	"fmt"
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddr          = common.EnvString("HTTP_ADDR", ":3000")
	ordersServiceAddr = "localhost:2000"
)

func main() {
	conn, err := grpc.NewClient(ordersServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dail server at: %v", err)
	}
	defer conn.Close()
	fmt.Println("Dialing orders service at", ordersServiceAddr)
	c := api.NewOrderServiceClient(conn)

	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoutes(mux)

	fmt.Println("starting http server at:", httpAddr)
	log.Fatal(http.ListenAndServe(httpAddr, mux))
}
