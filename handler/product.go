package handler

import (
	"TJMart/models"
	"encoding/json"
	"net/http"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("search")
	products := models.GetProducts()

	if keyword != "" {
		var filtered []models.Product
		//Sequential Search
		for _, p := range products {
			if p.Name == keyword {
				filtered = append(filtered, p)
			}
		}
		products = filtered
	}
	json.NewEncoder(w).Encode(products)
}
