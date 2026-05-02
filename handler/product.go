package handler

import (
	"TJMart/algorithm"
	"TJMart/models"
	"encoding/json"
	"net/http"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("search")
	products := models.GetProducts()

	if keyword != "" {
		products = algorithm.SequentialSearch(keyword, products, "name")
	}
	json.NewEncoder(w).Encode(products)
}
