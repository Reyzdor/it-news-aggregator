package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"trendfeed/internal/parser"
)

func main() {
	articles, err := parser.Pars()
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		jsonData, err := json.MarshalIndent(articles, "", " ")
		if err != nil {
			http.Error(w, "Failed", http.StatusInternalServerError)
			return
		}

		w.Write(jsonData)
	})

	http.ListenAndServe(":8080", nil)

}
