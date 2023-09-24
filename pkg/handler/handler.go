package handler

import (
	"github.com/dariaaal/golang-task2/internal"
	"github.com/dariaaal/golang-task2/pkg/service"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() {
	router := chi.NewRouter()

	logger, _ := zap.NewDevelopment()

	productTransport := internal.NewProductTransport(logger)

	router.Get("/products", internal.GetProducts)
	router.Post("/products", internal.AddProduct)
	router.Route("/product", func(r chi.Router) {
		r.Route("/{productId}", func(r chi.Router) {
			r.Use(productTransport.ProductCtx)
			r.Get("/", productTransport.GetProduct)
			r.Put("/", internal.UpdateProduct)
			r.Delete("/", internal.DeleteProduct)
		})
	})

}