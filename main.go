package main

import (
	"TJMart/handler"
	"fmt"
	"net/http"
)

func main() {
	port := "8080"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/login.html")
	})
	//api
	http.HandleFunc("/api/products", handler.GetAllProducts)
	http.HandleFunc("/api/login", handler.Login)
	fmt.Println("Server berjalan di http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
