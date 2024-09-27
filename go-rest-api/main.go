package main

import (
	"encoding/json"
	"fmt"
	"go-rest-api/database"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Item - struct to hold an individual item.
type Item struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var items []Item

// getItems responds with the list of all items as JSON.
func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// getItem responds with the details of a single item by ID.
func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get parameters
	for _, item := range items {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Item{})
}

// createItem adds a new item to the slice.
func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID = strconv.Itoa(len(items) + 1) // Mock ID - not safe for production!
	items = append(items, item)
	json.NewEncoder(w).Encode(item)
}

// main function to boot up everything
func main() {

	database.InitDB()

	r := mux.NewRouter()
	/* items = append(items, Item{ID: "1", Name: "Item One", Price: 10.99})
	items = append(items, Item{ID: "2", Name: "Item Two", Price: 20.99})

	r.HandleFunc("/api/items", getItems).Methods("GET")
	r.HandleFunc("/api/items/{id}", getItem).Methods("GET")
	r.HandleFunc("/api/items", createItem).Methods("POST") */

	r.HandleFunc("/api/item-create", database.CreateProdInDB).Methods("POST")
	r.HandleFunc("/api/items-db", database.GetItemsDB).Methods("GET")
	r.HandleFunc("/api/item-db/{id}", database.GetItemDB).Methods("GET")

	fmt.Println("Statring server...")
	log.Fatal(http.ListenAndServe(":8084", r))
}
