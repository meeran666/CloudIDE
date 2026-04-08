package routes

import (
	"encoding/json"
	"frontend/components"
	"frontend/models"
	"net/http"
)

func IDEHandler(w http.ResponseWriter, r *http.Request) {
	models.DirprofileArr = nil
	resp, err := http.Get("http://localhost:3000/")

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&models.DirprofileArr)
	if err != nil {
		panic(err)
	}
	components.IDEBase(models.DirprofileArr).Render(r.Context(), w)
	models.DirprofileArr = nil

}
