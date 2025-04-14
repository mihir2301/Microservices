package main

import (
	"common"
	"fmt"
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
)

var (
	httpAddr = common.EnvString("HTTP_ADDR", ":3000")
)

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)

	fmt.Println("starting http server at:", httpAddr)
	log.Fatal(http.ListenAndServe(httpAddr, mux))
}
