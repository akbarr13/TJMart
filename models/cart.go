package models

import (
	"TJMart/models"
	"encoding/json"
	"os"
)

type cart struct {
	Product  models.Product `json:"product"`
	Username models.User    `json:"username"`
}

const cartFile = "data/carts.json"

func getCart(name string) []cart {
	return []cart{}
}

func AddToCart(item models.Product, user models.User) bool {
	var cartItem = cart{Product: item, Username: user}
	byteValue, err := json.Marshal(cartItem)
	if err != nil {
		return false
	}

	err = os.WriteFile(cartFile, byteValue, 0644)
	return true
}
