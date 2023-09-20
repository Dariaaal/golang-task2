package internal

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func GetRouter() {

	logger, _ := zap.NewDevelopment()

	logger.Debug("starting server")

	productTransport := newProductTransport(logger)

	router := chi.NewRouter()
	// router.Get("/product/{id}", productTransport.GetProduct)
	router.Route("/product", func(r chi.Router) {
		r.Route("/{productId}", func(r chi.Router) {
			r.Use(productTransport.ProductCtx)
			r.Get("/", productTransport.GetProduct)
		})
	})
	// router.Use(ProductCtx).Get("/", productTransport.GetProduct)
	// router.Put("/product/{id}", putHandler)
	// router.Delete("/product/{id}", deleteHandler)
	port := ":8000"

	logger.Debug("server is listening", zap.String("port", port))

	err := http.ListenAndServe(port, router)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
