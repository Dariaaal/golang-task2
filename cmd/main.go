package main

import (
	"net/http"

	"github.com/dariaaal/golang-task2/internal/domain/product"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func main() {

	logger, _ := zap.NewDevelopment()

	logger.Debug("starting server")

	router := chi.NewRouter()
	router.Route("/product", internal.GetRoutes)

	port := ":8000"

	logger.Debug("server is listening", zap.String("port", port))

	err := http.ListenAndServe(port, router)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
