package handler

import (
	"TJMart/models"
	"encoding/json"
	"net/http"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products := models.GetProducts()
	json.NewEncoder(w).Encode(products)
}
