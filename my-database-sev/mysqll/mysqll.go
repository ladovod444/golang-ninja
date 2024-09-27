package mysqll

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	//"os"

	_ "github.com/go-sql-driver/mysql"
	//_ "github.com/lib/pq"
)

/* const (
	host = "localhost"
	//port     = 5432
	port     = 54322
	user     = "postgres"
	password = "postgres"
	dbname   = "db"
) */

const (
	host = "localhost"
	//port     = 5432
	port     = 33062
	user     = "drupal"
	password = "drupal"
	//dbname   = "db"
	dbname = "wallets"
)

var db *sql.DB

func initDB() {
	/* psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname) */

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		user, password, host, port, dbname)

	// db, err := sql.Open("mysql", "store_db_user:EXAMPLE_PASSWORD@tcp(127.0.0.1:3306)/store_db")
	db, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	/* var err error
	//db, err = sql.Open("postgres", psqlInfo)
	db, err = sql.Open("mysql", psqlInfo)
	if err != nil {
		log.Fatal(err)
	} */

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to!", dbname)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	//rows, err := db.Query("SELECT * FROM products")
	//rows, err := db.Query("SELECT ID, post_title FROM wp_posts where post_type = 'product'")
	rows, err := db.Query("SELECT ID, post_title FROM wp_posts")
	if err != nil {
		//log.Fatal("Errrr", err)
		panic(err.Error())
	} else {
		//fmt.Println(rows)
		fmt.Fprintf(w, "<p>Product with id = %v\n", rows)
	}
	//os.Exit(0)
	defer rows.Close()

	//products := make([]string, 0)
	products_map := make(map[int]string, 0)
	for rows.Next() {
		var ID int
		var post_title string
		//if err := rows.Scan(&id, &name); err != nil {
		if err := rows.Scan(&ID, &post_title); err != nil {
			log.Fatal("DDD", err)
		}
		//products = append(products, name)
		products_map[ID] = post_title
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
	rows, err := db.Query("INSERT INTO products (name) VALUES ('Galaxy S27')")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Fprintf(w, "<p>%v</p>", rows)
	}
	defer rows.Close()
}

func Mysqll() {
	initDB()

	http.HandleFunc("/products", getProducts)
	http.HandleFunc("/create-product", CreateProdInDB)
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		//panic("Some panic !!!", err)
		log.Fatal("Errrr", err)
	}
	println("Running code after ListenAndServe (only happens when server shuts down)")
}
