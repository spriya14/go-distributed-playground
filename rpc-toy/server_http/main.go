package main

import (
	"encoding/json"
	"goURL-shortie/rpc-toy/common"
	"log"
	"net/http"
)

func addHandler(w http.ResponseWriter, r *http.Request) {

	// Only allow POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Decoding request JSON into Args struct
	var args common.Args
	if err := json.NewDecoder(r.Body).Decode(&args); err != nil {
		http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
		return
	}
	// Perform Add operation
	result := args.A + args.B + len(args.Payload)

	// Prepare JSON response
	response := common.JsonResponse{Result: result}

	// Encode the response as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
func main() {
	http.HandleFunc("/add", addHandler)
	log.Println("HTTP server listening on port 9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
