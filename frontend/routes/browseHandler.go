package routes

import (
	"encoding/json"
	"fmt"
	"frontend/components"
	"frontend/models"
	"net/http"
)

func BrowseHandler(w http.ResponseWriter, r *http.Request) {

	path := r.FormValue("path")
	golet_id := r.FormValue("golet_id")
	//you have to find the golet id in backend by database
	//production part
	// golet_id := "weber"
	targetURL := "http://" + golet_id + ".localhost:3006/browse" + "?path=" + path
	//dev part
	// targetURL := "http://localhost:3003/browse?path=" + path
	resp, err := http.Get(targetURL)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 400)

	}
	defer resp.Body.Close()

	fmt.Println(models.DirprofileArr)
	err = json.NewDecoder(resp.Body).Decode(&models.DirprofileArr)
	fmt.Println(models.DirprofileArr)

	if err != nil {
		fmt.Println(err)
		// do a response of err to frontend in future
	}

	components.FileStructure(path, models.DirprofileArr, golet_id).Render(r.Context(), w)
	models.DirprofileArr = nil
}
