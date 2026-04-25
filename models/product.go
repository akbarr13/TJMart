package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
	Category string  `json:"Category"`
}

const productFile = "data/products.json"

func GetProducts() []Product {
	file, err := os.ReadFile(productFile)
	if err != nil {
		return []Product{}
	}

	var products []Product
	json.Unmarshal(file, &products)
	fmt.Print(products)
	return products
}
