package main

import (
	"fmt"
	"net/http"

	"github.com/antnose/Ecommerce/global_router"
	"github.com/antnose/Ecommerce/handlers"
	"github.com/antnose/Ecommerce/product"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))

	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProducts))

	fmt.Println("Server running on port: 3000")

	globalRouter := global_router.GlobalRouter(mux)

	err := http.ListenAndServe(":3000", globalRouter)

	if err != nil {
		fmt.Println("Error starting server", err)
	}
}

func init() {
	pd1 := product.Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is red",
		Price:       120,
		ImgURL:      "http://orange.com",
	}

	pd2 := product.Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is green",
		Price:       120,
		ImgURL:      "http://apple.com",
	}

	pd3 := product.Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is yellow",
		Price:       120,
		ImgURL:      "http://banana.com",
	}

	product.ProductList = append(product.ProductList, pd1, pd2, pd3)
}
