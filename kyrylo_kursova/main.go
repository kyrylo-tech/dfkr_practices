package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	connections []string
	mu          sync.Mutex
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		result := strings.Join(connections, ", ")
		mu.Unlock()

		w.Write([]byte(result))
	})

	r.Post("/connect", func(w http.ResponseWriter, r *http.Request) {
		var data map[string]any

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		val, ok := data["name"].(string)
		if !ok || val == "" {
			http.Error(w, "Name is required", http.StatusBadRequest)
			return
		}

		userName := strings.TrimSpace(val)

		mu.Lock()
		connections = append(connections, userName)
		mu.Unlock()

		fmt.Println("Added:", userName)
		w.Write([]byte("You are successfully added to session..."))
	})

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", r)
}
