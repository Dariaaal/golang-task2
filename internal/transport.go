package internal

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	// parse id
	// find product by id (for)
	// if no - 404, if yes - 200 and get product
	foundProduct, err := json.Marshal(products)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(foundProduct)
}
