package routes

import (
	"encoding/json"
	"net/http"
	"os"
	"runner/models"

	"github.com/fatih/color"
)

func FileReadHandler(w http.ResponseWriter, r *http.Request) {

	var req models.FileRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		color.Red("Error: %v", err)
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	//this is for dev
	// file_path := "../user_environment/" + req.Path

	//this is for prod
	file_path := req.Path
	data, err := os.ReadFile(file_path)
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write(data)
}
