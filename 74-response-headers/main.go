package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)

		_, err := fmt.Fprintln(w, "Resource created")
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
