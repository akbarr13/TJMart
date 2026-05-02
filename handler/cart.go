package handler

import (
	"TJMart/algorithm"
	"TJMart/models"
	"encoding/json"
	"net/http"
)

func AddCart(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("productId")
	products = algorithm.SequentialSearch(productId, products, "id")
	json.NewEncoder(w).Encode(products)
}
