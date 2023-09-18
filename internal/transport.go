package internal

import (
	"encoding/json"

	"net/http"

	"go.uber.org/zap"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	// parse id
	// find product by id (for)
	// if no - 404, if yes - 200 and get product

	// vars := mux.products(r)
	// key := vars["uid"]

	// parsed := json.Unmarshal([]byte{}, products)
	// key := parsed["uid"]

	// productsArr := []Product{}
	// parsed := json.Unmarshal([]byte(products), &productsArr)
	// key := productsArr["uid"]

	for _, product := range products {
		if product.ID == key {
			json.NewEncoder(w).Encode(product)
		}
	}

	logger, _ := zap.NewProduction()

	foundProduct, err := json.Marshal(products)
	if err != nil {
		logger.Fatal(err.Error())
	}
	w.Write(foundProduct)
}
