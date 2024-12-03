package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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
	handlers    map[string]func(http.ResponseWriter, *http.Request)
	middlewares []func(func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request)
}

func NewAppServer() *AppServer {
	return &AppServer{
		handlers: make(map[string]func(http.ResponseWriter, *http.Request)),
	}
}

// to add the handlers
func (appServer *AppServer) Add(resource string, handlerFn func(http.ResponseWriter, *http.Request)) {
	for i := len(appServer.middlewares) - 1; i >= 0; i-- {
		middleware := appServer.middlewares[i]
		handlerFn = middleware(handlerFn)
	}
	appServer.handlers[resource] = handlerFn
}

func (appServer *AppServer) Use(middleware func(func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request)) {
	appServer.middlewares = append(appServer.middlewares, middleware)
}

func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler := appServer.handlers[r.URL.Path]; handler != nil {
		handler(w, r)
		return
	}
	http.Error(w, "resource not found", http.StatusNotFound)
}

// app specific handlers
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// simulating a time consuming operation
	fmt.Println("[indexHandler] trace id :", r.Context().Value("trace-id"))
	time.Sleep(1 * time.Second)
	select {
	case <-r.Context().Done():
		return
	default:
		fmt.Fprintln(w, "Hello, World!")
	}

}

func productHandler(w http.ResponseWriter, r *http.Request) {
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
}

func customersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "The list of customers will be served")
}

// middlewares
func logMiddleware(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s - %s\n", r.Method, r.URL.Path)
		handler(w, r)
	}
}

func timeoutMiddleware(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		timeoutCtx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		handler(w, r.WithContext(timeoutCtx))
		if timeoutCtx.Err() != nil {
			fmt.Println(timeoutCtx.Err())
		}
	}
}

func traceMiddleware(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		traceIdCtx := context.WithValue(r.Context(), "trace-id", 100)
		handler(w, r.WithContext(traceIdCtx))
	}
}

func main() {
	appServer := NewAppServer()
	appServer.Use(traceMiddleware)
	appServer.Use(timeoutMiddleware)
	appServer.Use(logMiddleware)
	appServer.Add("/", indexHandler)
	appServer.Add("/products", productHandler)
	appServer.Add("/customers", customersHandler)
	http.ListenAndServe(":8080", appServer)
}
