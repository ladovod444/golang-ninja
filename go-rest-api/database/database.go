package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	//port     = 5432
	port     = 54322
	user     = "postgres"
	password = "postgres"
	dbname   = "db"
)

var db *sql.DB

const port_value string = "8084"

type Item struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ItemRes struct {
	ID    int  
	Name  string  
	Price float64 
}

var items []Item

func InitDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")
	/* fmt.Printf("Use http://localhost:%s/products in you browser to list products\n", port_value)
	fmt.Printf("Use http://localhost:%s/create-product?product=[PRODUCT_NAME] to create a test\n", port_value) */
}

func GetItemsDB(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//products := make([]string, 0)
	//products_map := make(map[int]string, 0)
	products_map := make(map[int]ItemRes, 0)
	for rows.Next() {
		var id int
		var name string
		var price float64
		if err := rows.Scan(&id, &name, &price); err != nil {
			log.Fatal("DDD", err)
		}
		//products_map[id] = name
		itemRes := ItemRes{id, name, price}
		products_map[id] = itemRes
	}

	// Output resultS in json
	json.NewEncoder(w).Encode(products_map)
	//json.NewEncoder(w).Encode([]int{2, 3, 4, 5, 6})
}

func GetItemDB(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// GET id (e.g. 16) from url http://localhost:8084/api/item-db/16
	params := mux.Vars(r)

	select_string := fmt.Sprintf("SELECT * FROM products WHERE id=%s", params["id"])

	rows, err := db.Query(select_string)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//products := make([]string, 0)
	//products_map := make(map[int]string, 0)
	products_map := make(map[int]ItemRes, 0)
	for rows.Next() {
		var id int
		var name string
		var price float64
		if err := rows.Scan(&id, &name, &price); err != nil {
			log.Fatal("DDD", err)
		}
		//products = append(products, name)
		//products_map[id] = name
		itemRes := ItemRes{id, name, price}
		products_map[id] = itemRes
	}

	// Output result in json
	json.NewEncoder(w).Encode(products_map)
	//json.NewEncoder(w).Encode([]int{2, 3, 4, 5, 6})
}

func CreateProdInDB(w http.ResponseWriter, r *http.Request) {
	//rows, err := db.Query("INSERT INTO products VALUES (3, 'Galaxy S9')")

	//product_name := r.URL.Query().Get("product")
	w.Header().Set("Content-Type", "application/json")
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID = strconv.Itoa(len(items) + 1) // Mock ID - not safe for production!
	if item.Name == "" {
		log.Fatal("Product name is not specified")
	} else {
		var lastInsertId int

		err := db.QueryRow("INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id", item.Name, item.Price).Scan(&lastInsertId)
	
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "Inserted product %s, price=%f with ID: %d\n", item.Name, item.Price, lastInsertId)

	}

}
