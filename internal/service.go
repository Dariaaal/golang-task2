package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	// json.NewEncoder(w).Encode(product)
	js, err := json.Marshal(product)
    if err != nil{
        log.Fatal(err)
	}
    w.Write(js)
}