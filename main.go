package main

import (
	authpb "api-gateway/protocol-buffer/golang/auth"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net/http"
	"os"
)

func init() {
	err := 	godotenv.Load()
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	if err != nil {
		log.Fatal(err)
	}

	log.SetFormatter(&log.TextFormatter{
		DisableQuote: true,
	})
}

func main() {
	ctx := context.Background()
	mux := runtime.NewServeMux()

	options := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	if err := authpb.RegisterAuthHandlerFromEndpoint(
		ctx,
		mux,
		os.Getenv("AUTH_ADDR"),
		options); err != nil {
		log.Fatalf("failed to register gRPC gateway %s", err)
	}

	log.Printf("start HTTP server on %s port", os.Getenv("PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), mux); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

