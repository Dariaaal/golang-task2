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

func (t *ProductTransport) ProductCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		productID := chi.URLParam(r, "productId")
		t.log.Debug("product id", zap.String("id", productID))
		ctx := context.WithValue(r.Context(), "id", productID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	foundProducts, err := json.Marshal(products)

	if err != nil {
		http.Error(w, http.StatusText(400), 400)
	}

	w.Write(foundProducts)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
    reqBody, _ := ioutil.ReadAll(r.Body)

	var newProduct *Product

	json.Unmarshal(reqBody, &newProduct)
	products = append(products, newProduct)

	json.NewEncoder(w).Encode(newProduct)
}

func (t *ProductTransport) GetProduct(w http.ResponseWriter, r *http.Request) {

	rowProductID := "2"

	// if rowProductID == nil {
	// 	http.Error(w, http.StatusText(400), 400)
	// 	return
	// }

	// productId := rowProductID.(string)

	t.log.Debug("get product", zap.String("id", rowProductID))

	for _, product := range products {
		if product.ID == rowProductID {
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

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
    id := "3"

    for index, product := range products {
        if product.ID == id {
            products = append(products[:index], products[index+1:]...)
        }
    }
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := "2"

	reqBody, _ := ioutil.ReadAll(r.Body)

	var newProduct *Product
	json.Unmarshal(reqBody, &newProduct)

	for index, product := range products {
        if product.ID == id {
            product.Cover = newProduct.Cover
			product.Title = newProduct.Title
			product.Description = newProduct.Description
			product.Price = newProduct.Price

			products[index] = product
        }
    }
}
