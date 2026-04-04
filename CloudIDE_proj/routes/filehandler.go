package routes

import (
	"encoding/json"
	"net/http"
	"os"
	"project/models"

	"github.com/fatih/color"
)

func FileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req models.FileRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		color.Red("Error: %v", err)
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	data, err := os.ReadFile(req.Path)
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write(data)
}
