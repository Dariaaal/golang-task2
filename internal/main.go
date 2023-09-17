package main

import (
	"fmt"

	"github.com/go-chi/chi/v5"
	// "go.uber.org/zap"
	// getHandler "./internal/service.go"
)

func main() {
	fmt.Println("Starting server...")

	router := chi.NewRouter()
	router.Get("/product/{id}", getHandler)
	// router.Put("/product/{id}", putHandler)
	// router.Delete("/product/{id}", deleteHandler)
	port := ":8000"

	// logger, _ := zap.NewProduction()

	// logger.Fatal(http.ListenAndServe(port, router))
	fmt.Printf("Server is listening on port %s...", port)
}

