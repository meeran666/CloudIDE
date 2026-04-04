package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runner/models"

	"github.com/fatih/color"
)

func FileHandler(w http.ResponseWriter, r *http.Request) {

	// CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	//  Handle OPTIONS (preflight)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	//  Now POST will come here
	if r.Method == http.MethodPost {
		fmt.Println("POST received")
	}
	var req models.FileRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		color.Red("Error: %v", err)
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	file_path := "../user_environment/" + req.Path
	data, err := os.ReadFile(file_path)
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write(data)
}
