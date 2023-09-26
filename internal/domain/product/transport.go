package internal

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type ProductTransport struct {
	log  *zap.Logger
	repo ProductStorage
}

func NewProductTransport(log *zap.Logger) ProductTransport {
	return ProductTransport{
		log:  log,
		repo: NewInMemoryStorage(),
	}
}

func GetRoutes(router chi.Router) {
	logger, _ := zap.NewDevelopment()

	productTransport := NewProductTransport(logger)

	router.Get("/all", productTransport.GetProducts)
	router.Post("/", productTransport.AddProduct)
	router.Route("/{productId}", func(r chi.Router) {
		r.Use(productTransport.ProductCtx)
		r.Get("/", productTransport.GetProduct)
		// r.Put("/", productTransport.UpdateProduct)
		r.Delete("/", productTransport.DeleteProduct)
	})

}

func (t *ProductTransport) ProductCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		productID := chi.URLParam(r, "productId")
		t.log.Debug("product id", zap.String("id", productID))
		ctx := context.WithValue(r.Context(), "id", productID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (t *ProductTransport) GetProducts(w http.ResponseWriter, r *http.Request) {
	foundProducts, err := json.Marshal(t.repo.GetAll())

	if err != nil {
		http.Error(w, http.StatusText(400), 400)
	}

	w.Write(foundProducts)
}

func (t *ProductTransport) AddProduct(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newProduct *Product
	json.Unmarshal(reqBody, &newProduct)
	err := NewService(t.repo).Add(*newProduct)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (t *ProductTransport) GetProduct(w http.ResponseWriter, r *http.Request) {

	productId := NewService(t.repo).GetById(r.Context()).ID

	t.log.Debug("get product", zap.String("id", productId))

	responseData, err := json.Marshal(NewService(t.repo).GetById(r.Context()))
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}

func (t *ProductTransport) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	err := NewService(t.repo).Delete(r.Context())
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// func (t *ProductTransport) UpdateProduct(w http.ResponseWriter, r *http.Request) {
// 	rowProductID := r.Context().Value("id")

// 	if rowProductID == nil {
// 		http.Error(w, http.StatusText(400), 400)
// 		return
// 	}

// 	productId := rowProductID.(string)

// 	reqBody, _ := ioutil.ReadAll(r.Body)

// 	var newProduct *Product
// 	json.Unmarshal(reqBody, &newProduct)

// 	for index, product := range products {
// 		if product.ID == productId {
// 			product.Cover = newProduct.Cover
// 			product.Title = newProduct.Title
// 			product.Description = newProduct.Description
// 			product.Price = newProduct.Price

// 			products[index] = product
// 		}
// 	}
// }
