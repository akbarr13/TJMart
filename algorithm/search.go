package algorithm

import (
	"TJMart/models"
)

func SequentialSearch(keyword string, products []models.Product, searchBy string) []models.Product {
	var filtered []models.Product

	for _, p := range products {
		switch searchBy {
		case "name":
			if p.Name == keyword {
				filtered = append(filtered, p)
			}
		case "id":
			if string(p.ID) == keyword {
				filtered = append(filtered, p)
			}
		}

	}
	return filtered
}
