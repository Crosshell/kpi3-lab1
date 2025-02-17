package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time string `json:"time"`
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	currentTime := time.Now().Format(time.RFC3339)
	response := TimeResponse{Time: currentTime}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Printf("Error encoding response: %v", err)
	}
}

func main() {
	http.HandleFunc("/time", timeHandler)

	log.Println("Starting server on :8795...")
	if err := http.ListenAndServe(":8795", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

//fdssfsdfdfs
