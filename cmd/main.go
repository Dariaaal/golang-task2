package main

import (
	"net/http"

	"github.com/dariaaal/golang-task2/internal/domain/product"
	"go.uber.org/zap"
)

func main() {

	logger, _ := zap.NewDevelopment()

	logger.Debug("starting server")

	// internal.GetRoutes()

	// create handler and use productTransport.GetRoutes()

	port := ":8000"

	logger.Debug("server is listening", zap.String("port", port))

	err := http.ListenAndServe(port, internal.GetRoutes())
	if err != nil {
		logger.Fatal(err.Error())
	}
}
