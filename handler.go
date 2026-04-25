package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, `
				<html>
					<body>
						<h1>Halo ini website golang!!!</h1>
					</body
				</html
				`)
	case "/welcome":
		fmt.Fprint(w, `
				<html>
					<body>
						<h1>SElamat datang!!!</h1>
					</body
				</html>
				`)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `
				<html>
					<body>
						<h1>Error bos!!!!</h1>
					</body
				</html>
				`)

	}
}
