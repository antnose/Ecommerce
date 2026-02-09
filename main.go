package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imgUrl"`
}

var productList []Product

func Greetings(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to application")
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request. Please make a GET request", http.StatusBadRequest)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(productList)

}

func CreateProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request. Please make a POST request", http.StatusBadRequest)
		return
	}

	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	encoder := json.NewEncoder(w)
	encoder.Encode(newProduct)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/greet", Greetings)

	mux.HandleFunc("/products", GetProducts)

	mux.HandleFunc("/create-products", CreateProducts)

	fmt.Println("Server running on port: 3000")
	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		fmt.Println("Error starting server", err)
	}
}

func init() {
	pd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is red",
		Price:       120,
		ImgURL:      "http://orange.com",
	}

	pd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is green",
		Price:       120,
		ImgURL:      "http://apple.com",
	}

	pd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is yellow",
		Price:       120,
		ImgURL:      "http://banana.com",
	}

	productList = append(productList, pd1, pd2, pd3)
}
