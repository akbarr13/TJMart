package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := "8080"
	http.HandleFunc("/", handler)
	fmt.Println("Server berjalan di http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
