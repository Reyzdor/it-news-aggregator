package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
	"trendfeed/internal/parser"
)

var (
	articles []parser.Article
	mu       sync.RWMutex
)

func updateNews() {
	data, err := parser.Pars()
	if err != nil {
		log.Println("Parse error", err)
		return
	}

	mu.Lock()
	articles = data
	mu.Unlock()
}

func startUpdater() {
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			updateNews()
		}
	}()
}

func main() {
	updateNews()

	startUpdater()

	http.HandleFunc("/api/news", func(w http.ResponseWriter, r *http.Request) {
		mu.RLock()
		defer mu.RUnlock()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(articles)
	})

	fs := http.FileServer(http.Dir("./client/static"))
	http.Handle("/", fs)

	http.ListenAndServe(":8080", nil)

}
