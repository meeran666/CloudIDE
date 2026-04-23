package routes

import (
	"frontend/helpers"
	"frontend/models"
	"log"
	"net/http"
)

func ChangeHandler(w http.ResponseWriter, r *http.Request) {
	golet_id := r.FormValue("golet_id")
	//dev part
	// targetURL := "http://localhost:3003"

	//prod part
	targetURL := "http://" + golet_id + ".localhost:3006"

	log.Printf("Proxying: %s /save-change → localhost:3006/save-change", r.Method)
	models.Proxy = nil

	helpers.Proxy(golet_id, targetURL)

	models.Proxy.ServeHTTP(w, r)
}
