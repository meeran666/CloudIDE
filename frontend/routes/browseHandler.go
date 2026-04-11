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
	//you have to find the golet id in backend by database
	golet_id := "service"
	fmt.Println("path")
	fmt.Println(path)
	// 2. Target URL
	apiUrl := "http://" + golet_id + ".localhost:3006/browse" + "?path=" + path
	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 400)

	}

	defer resp.Body.Close()
	fmt.Println("Status:", resp.Status)

	// resp, err = http.Get("http://localhost:3000/browse")
	// if err != nil {
	// 	fmt.Println(err)
	// 	http.Error(w, err.Error(), 400)
	// }
	// defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&models.DirprofileArr)
	if err != nil {
		fmt.Println(err)
		// do a response of err to frontend in future
	}

	components.FileStructure(path, models.DirprofileArr).Render(r.Context(), w)
	models.DirprofileArr = nil
}
