package main

import (
    "encoding/json"
    "log"
    "net/http"
    //"strconv"
    "github.com/gorilla/mux"
)

type Product struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
}

var products []Product
var nextID = 1

func getProducts(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

	//products = append(products, Product{1, "new prod"})
	//products = append(products, Product{2, "new prod2"})
	//products = append(products, "new prod2")

    json.NewEncoder(w).Encode(products)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
    var product Product
    if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    product.ID = nextID
    nextID++
    products = append(products, product)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(product)
}

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/products", getProducts).Methods("GET")
    router.HandleFunc("/products", createProduct).Methods("POST")

    log.Fatal(http.ListenAndServe(":8080", router))
}