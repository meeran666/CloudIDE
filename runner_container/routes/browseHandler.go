package routes

import (
	"encoding/json"
	"net/http"
	"runner/models"

	"github.com/fatih/color"
)

func BrowseHandler(w http.ResponseWriter, r *http.Request) {

	path := r.FormValue("path")
	err := filelist(path)

	if err != nil {
		color.Red("Error: %v", err)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	// Set response type
	w.Header().Set("Content-Type", "application/json")

	// Encode and return JSON
	json.NewEncoder(w).Encode(models.DirprofileArr)
	models.DirprofileArr = nil
}
