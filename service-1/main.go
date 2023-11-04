package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handleGet(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path)

	if r.URL.Path != "/" {
		http.Error(w, "Page not found", http.StatusNotFound)
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	payload := make(map[string] string)
	payload["message"] = "Hello from service-1"
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleGet)

	log.Println("Launching server on port 8080...")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal("ListenAndServe:", err)
}
