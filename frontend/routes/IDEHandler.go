package routes

import (
	"encoding/json"
	"fmt"
	"frontend/components"
	"frontend/models"
	"net/http"
)

func IDEHandler(w http.ResponseWriter, r *http.Request) {
	golet_id := r.URL.Query().Get("golet_id")
	golet_id = "service"
	apiUrl := "http://" + golet_id + ".localhost:3006"
	resp, err := http.Get(apiUrl)

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	models.DirprofileArr = nil
	err = json.NewDecoder(resp.Body).Decode(&models.DirprofileArr)
	if err != nil {
		fmt.Println(err)
	}

	// components.IDEBase(models.DirprofileArr).Render(r.Context(), w)
	components.Base(components.IDEBase(models.DirprofileArr)).Render(r.Context(), w)

	models.DirprofileArr = nil

}
