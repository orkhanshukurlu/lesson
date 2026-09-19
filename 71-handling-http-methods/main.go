package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			_, err := fmt.Fprint(w, "GET request received")
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		if r.Method == http.MethodPost {
			_, err := fmt.Fprint(w, "POST request received")
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
