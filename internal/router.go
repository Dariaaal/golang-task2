package internal

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func GetRouter() {

	logger, _ := zap.NewDevelopment()

	logger.Debug("starting server")

	productTransport := NewProductTransport(logger)

	router := chi.NewRouter()

	router.Get("/products", GetProducts)
	router.Post("/products", AddProduct)
	router.Route("/product", func(r chi.Router) {
		r.Route("/{productId}", func(r chi.Router) {
			r.Use(productTransport.ProductCtx)
			r.Get("/", productTransport.GetProduct)
			r.Put("/", UpdateProduct)
			r.Delete("/", DeleteProduct)
		})
	})

	port := ":8000"

	logger.Debug("server is listening", zap.String("port", port))

	err := http.ListenAndServe(port, router)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
