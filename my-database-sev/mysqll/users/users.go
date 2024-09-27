package users

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host1 = "localhost"
	//port     = 5432
	port1     = 33062
	user1     = "drupal"
	password1 = "drupal"
	//dbname   = "db"
	//dbname1 = "exampledb"
	dbname1 = "wallets"
)

func CreateUserInDB(w http.ResponseWriter, r *http.Request) {
	// db, err := sql.Open("mysql", "user:password@/exampledb")
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		user1, password1, host1, port1, dbname1)

	// db, err := sql.Open("mysql", "store_db_user:EXAMPLE_PASSWORD@tcp(127.0.0.1:3306)/store_db")
	db, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ensure connection is established
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")
	fmt.Println("From URL:", r.URL.Path)
	uname := r.URL.Query().Get("uname")
	age_par := r.URL.Query().Get("age")
	//fmt.Println("uname From URL:", uname)
	if uname == "" || age_par == "" {
		log.Fatal("username or age are not specified")
	} else {

		// Insert data
		res, err := db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", uname, age_par)
		if err != nil {
			log.Fatal(err)
		}
		lastId, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "Inserted user with ID: %d\n", lastId)

		// Query data
		var (
			id   int
			name string
			age  int
		)
		rows, err := db.Query("SELECT id, name, age FROM users")
		//rows, err := db.Query("SELECT ID, post_title FROM wp_posts")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id, &name, &age)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprintf(w, "%d: %s, %d\n", id, name, age)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}

}

func getProductsMy(w http.ResponseWriter, r *http.Request) {
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		user1, password1, host1, port1, dbname1)

	// db, err := sql.Open("mysql", "store_db_user:EXAMPLE_PASSWORD@tcp(127.0.0.1:3306)/store_db")
	db, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ensure connection is established
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Connected!\n")

	// Query data
	/* var (
		id   int
		name string
		age  int
	) */

	var (
		ID         int
		post_title string
		//age  int
	)
	//rows, err := db.Query("SELECT id, name, age FROM users")
	rows, err := db.Query("SELECT ID, post_title FROM wp_posts WHERE post_type='product'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	//products_map := make(map[int]string, 0)
	fmt.Fprintf(w, "Fetching products!\n")
	for rows.Next() {
		err := rows.Scan(&ID, &post_title)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("%d: %s, %d\n", id, name, age)
		//products_map[ID] = post_title
		//fmt.Fprintf(w, "<p>User with id = %d, name=%s, age=%d</p>", ID, post_title, age)
		fmt.Fprintf(w, "<p>Post with id = %d, title=%s, </p>\n", ID, post_title)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		user1, password1, host1, port1, dbname1)

	// db, err := sql.Open("mysql", "store_db_user:EXAMPLE_PASSWORD@tcp(127.0.0.1:3306)/store_db")
	db, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ensure connection is established
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Connected!\n")

	// Query data
	var (
		id   int
		name string
		email  string
	)


	//rows, err := db.Query("SELECT id, name, age FROM users")
	rows, err := db.Query("SELECT ID, user_nicename, user_email FROM wp_users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	//products_map := make(map[int]string, 0)
	fmt.Fprintf(w, "Fetching users!\n")
	for rows.Next() {
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			log.Fatal(err)
		}
		
		//products_map[ID] = post_title
		fmt.Fprintf(w, "<p>User with id = %d, name=%s, email=%s</p>\n", id, name, email)
		
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func Init() {
	http.HandleFunc("/products", getProductsMy)
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/create-user", CreateUserInDB)
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		//panic("Some panic !!!", err)
		log.Fatal("Errrr", err)
	}
}
