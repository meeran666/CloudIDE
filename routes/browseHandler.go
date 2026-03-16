package routes

import (
	"net/http"
	"project/components"
	"project/models"

	"github.com/fatih/color"
)

func BrowseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	path := r.FormValue("path")
	err := filelist(path)
	if err != nil {
		color.Red("Error: %v", err)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	components.FileStructure(path, models.DirprofileArr).Render(r.Context(), w)
	models.DirprofileArr = nil
}
