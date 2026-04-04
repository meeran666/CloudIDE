package routes

import (
	"encoding/json"
	"frontend/components"
	"frontend/models"
	"net/http"
)

func IDEHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
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
	components.Base(models.DirprofileArr).Render(r.Context(), w)
	models.DirprofileArr = nil

}
