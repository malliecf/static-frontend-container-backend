package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var widgets = []map[string]interface{}{
	{"id": 1, "name": "Widget A"},
	{"id": 2, "name": "Widget B"},
	{"id": 3, "name": "Widget C"},
}

func widgetsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/api/widgets")

	if path == "" || path == "/" {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		json.NewEncoder(w).Encode(widgets)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idString := strings.TrimPrefix(path, "/")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid widget ID", http.StatusBadRequest)
		return
	}

	for _, widget := range widgets {
		if widget["id"] == id {
			json.NewEncoder(w).Encode(widget)
			return
		}
	}

	http.Error(w, "Widget not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/api/widgets", widgetsHandler)
	http.HandleFunc("/api/widgets/", widgetsHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
