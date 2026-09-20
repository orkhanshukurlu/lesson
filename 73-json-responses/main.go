package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		user := User{
			Name:  "John Doe",
			Email: "john.doe@example.com",
			Age:   30,
		}

		w.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(w).Encode(user)
		if err != nil {
			http.Error(w, "Error has occurred", http.StatusInternalServerError)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
