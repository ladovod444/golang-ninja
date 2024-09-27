package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// User represents the user model
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var (
	users = []*User{}
	mutex = &sync.Mutex{}
)

func main() {
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/user/", userHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	case "POST":
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		mutex.Lock()
		users = append(users, &user)
		mutex.Unlock()
		json.NewEncoder(w).Encode(user)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/user/"):]
	mutex.Lock()
	defer mutex.Unlock()

	for _, user := range users {
		if user.ID == id {
			switch r.Method {
			case "GET":
				fmt.Println("Got user ", id)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(user)
			case "DELETE":
				for i, u := range users {
					if u.ID == id {
						users = append(users[:i], users[i+1:]...)
						fmt.Println("user deleted ", id)
						w.WriteHeader(http.StatusOK)
						return
					}
				}
			default:
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
			return
		}
	}

	http.NotFound(w, r)
}