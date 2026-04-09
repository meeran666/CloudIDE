package routes

import (
	"encoding/json"
	"frontend/components"
	"frontend/models"
	"net/http"
)

func BrowseHandler(w http.ResponseWriter, r *http.Request) {

	path := r.FormValue("path")
	resp, err := http.Get("http://localhost:3000/browse")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&models.DirprofileArr)
	if err != nil {
		panic(err)
	}
	components.FileStructure(path, models.DirprofileArr).Render(r.Context(), w)
	models.DirprofileArr = nil
}
