package main

import (
	"fmt"
	"net/http"
	"trendfeed/internal/parser"
)

func main() {
	parser.Pars()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Trend Feed!")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error")
	}
}
