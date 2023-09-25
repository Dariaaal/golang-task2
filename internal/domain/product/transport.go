package internal

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type ProductTransport struct {
	log *zap.Logger
}

func NewProductTransport(log *zap.Logger) ProductTransport {
	return ProductTransport{
		log: log,
	}
}

func GetRoutes() *chi.Mux {
	router := chi.NewRouter()

	logger, _ := zap.NewDevelopment()

	productTransport := NewProductTransport(logger)

	router.Route("/product", func(r chi.Router) {
		router.Get("/", productTransport.GetProduct)
		router.Post("/", productTransport.AddProduct)
		r.Route("/{productId}", func(r chi.Router) {
			r.Use(productTransport.ProductCtx)
			r.Get("/", productTransport.GetProduct)
			r.Put("/", productTransport.UpdateProduct)
			r.Delete("/", productTransport.DeleteProduct)
		})
	})

	return router
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
	foundProducts, err := json.Marshal(products)

	if err != nil {
		http.Error(w, http.StatusText(400), 400)
	}

	w.Write(foundProducts)
}

func (t *ProductTransport) AddProduct(w http.ResponseWriter, r *http.Request) {
    reqBody, _ := ioutil.ReadAll(r.Body)

	var newProduct *Product

	json.Unmarshal(reqBody, &newProduct)
	products = append(products, newProduct)

	json.NewEncoder(w).Encode(newProduct)
}

func (t *ProductTransport) GetProduct(w http.ResponseWriter, r *http.Request) {

	rowProductID := r.Context().Value("id")

	if rowProductID == nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	productId := rowProductID.(string)

	t.log.Debug("get product", zap.String("id", productId))

	for _, product := range products {
		if product.ID == productId {
			responseData, err := json.Marshal(product)
			if err != nil {
				http.Error(w, http.StatusText(400), 400)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(responseData)
			return
		}
	}

	http.Error(w, http.StatusText(400), 400)
}

func (t *ProductTransport) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	rowProductID := r.Context().Value("id")

	if rowProductID == nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	productId := rowProductID.(string)

    for index, product := range products {
        if product.ID == productId {
            products = append(products[:index], products[index+1:]...)
        }
    }
}

func (t *ProductTransport) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	rowProductID := r.Context().Value("id")

	if rowProductID == nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	productId := rowProductID.(string)

	reqBody, _ := ioutil.ReadAll(r.Body)

	var newProduct *Product
	json.Unmarshal(reqBody, &newProduct)

	for index, product := range products {
        if product.ID == productId {
            product.Cover = newProduct.Cover
			product.Title = newProduct.Title
			product.Description = newProduct.Description
			product.Price = newProduct.Price

			products[index] = product
        }
    }
}
