package main

import (
	// "encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func main() {
	fmt.Println("Starting server...")

	router := chi.NewRouter()
	router.Get("/product/{id}", getHandler)
	// router.Put("/product/{id}", putHandler)
	// router.Delete("/product/{id}", deleteHandler)
	port := ":8000"

	logger, _ := zap.NewProduction()

	logger.Error(http.ListenAndServe(port, router))
	fmt.Println("Server is listening on port 8000...")
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	// json.NewEncoder(w).Encode()
}
