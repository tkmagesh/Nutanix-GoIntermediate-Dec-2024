package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Model
type Product struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Cost     float64 `json:"cost"`
	Category string  `json:"category"`
}

// Inmemory data
var products []Product = []Product{
	Product{Id: 100, Name: "Pen", Cost: 10, Category: "stationary"},
	Product{Id: 101, Name: "Pencil", Cost: 5, Category: "stationary"},
	Product{Id: 102, Name: "Marker", Cost: 50, Category: "stationary"},
}

type AppServer struct {
}

func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s - %s\n", r.Method, r.URL.Path)
	switch r.URL.Path {
	case "/":
		fmt.Fprintln(w, "Hello, World!")
	case "/products":
		// fmt.Fprintln(w, "The list of products will be served")
		switch r.Method {
		case http.MethodGet:
			if err := json.NewEncoder(w).Encode(products); err != nil {
				http.Error(w, "error encoding data", http.StatusInternalServerError)
			}
		case http.MethodPost:
			var newProduct Product
			if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
				http.Error(w, "invalid payload", http.StatusBadRequest)
				return
			}
			products = append(products, newProduct)
			w.WriteHeader(http.StatusCreated)
			if err := json.NewEncoder(w).Encode(newProduct); err != nil {
				http.Error(w, "error encoding data", http.StatusInternalServerError)
			}
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}

	case "/customers":
		fmt.Fprintln(w, "The list of customers will be served")
	default:
		http.Error(w, "resource not found", http.StatusNotFound)
	}

}

func main() {
	appServer := &AppServer{}
	http.ListenAndServe(":8080", appServer)
}
