package main

import (
	"database/sql"
	"fmt"
	"log"
	"my-database-sev/mysqll/users"
	"net/http"
	"os"

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

const port_value string = "8082"

func initDB() {
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
	fmt.Printf("Use http://localhost:%s/products in you browser to list products\n", port_value)
	fmt.Printf("Use http://localhost:%s/create-product?product=[PRODUCT_NAME] to create a test\n", port_value)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//products := make([]string, 0)
	products_map := make(map[int]string, 0)
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal("DDD", err)
		}
		//products = append(products, name)
		products_map[id] = name
	}

	//fmt.Fprintf(w, "Products: %v", products)
	//fmt.Fprintf(w, "Products (map): %v", products_map)

	for id, name := range products_map {
		//fmt.Fprintf(w, "<p>Product with id = %d, name=%s</p>\n", id, name)
		fmt.Fprintf(w, "<p>Product with id = %d, name=%s</p>", id, name)
	}
}

// https://metanit.com/sql/postgresql/3.1.php
func CreateProdInDB(w http.ResponseWriter, r *http.Request) {
	//rows, err := db.Query("INSERT INTO products VALUES (3, 'Galaxy S9')")

	product_name := r.URL.Query().Get("product")
	//age_par := r.URL.Query().Get("age")
	if product_name == "" {
		log.Fatal("Product name is not specified")
	} else {

		//rows, err := db.Query("INSERT INTO products (name) VALUES (?)", product_name)
		var lastInsertId int
		// Insert data
		//db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", uname, age_par)
		//insertString := fmt.Sprintf("INSERT INTO products (id, name) VALUES (%d, '%s')", 10, product_name)
		//insertString := fmt.Sprintf("INSERT INTO products (name) VALUES ('%s')", product_name)
		//res, err := db.Exec("INSERT INTO products (id, name) VALUES (?, ?)", 10, product_name)
		//res, err := db.Exec(insertString).Scan(lastId)
		// FOR POSTGRESS USE this func to get last inserted id
		err := db.QueryRow("INSERT INTO products (name) VALUES ($1) RETURNING id", product_name).Scan(&lastInsertId)
		//rows, err := db.Query("INSERT INTO products VALUES (3, 'Galaxy S9')")
		//defer res.Close()
		if err != nil {
			log.Fatal(err)
		}
		//lastId, err := res.LastInsertId()
		/* if err != nil {
			log.Fatal(err)
		} */
		fmt.Fprintf(w, "Inserted product with ID: %d\n", lastInsertId)

		/* if err != nil {
			log.Fatal(err)
		} else {
			fmt.Fprintf(w, "<p>%v</p>", rows)
		} */
	}

}

func main() {

	//mysqll.Init()

	// FOR POSTGRES
	initDB()
	http.HandleFunc("/products", getProducts)
	http.HandleFunc("/create-product", CreateProdInDB)
	http.ListenAndServe(":"+port_value, nil)

	os.Exit(0)
	// FOR MYSQL
	users.Init()

}
