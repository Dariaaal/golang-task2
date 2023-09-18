package internal

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func GetRouter() {

	logger, _ := zap.NewProduction()

	logger.Debug("starting server")

	router := chi.NewRouter()
	router.Get("/product/{id}", GetProduct)
	// router.Put("/product/{id}", putHandler)
	// router.Delete("/product/{id}", deleteHandler)
	port := ":8000"

	logger.Debug("server is listening", zap.String("port", port))

	err := http.ListenAndServe(port, router)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
