package routes

import (
	"frontend/helpers"
	"frontend/models"
	"log"
	"net/http"
)

func DeleteServiceProxy(w http.ResponseWriter, r *http.Request) {
	golet_id := r.FormValue("golet_id")
	targetURL := "http://localhost:3003"

	log.Printf("Proxying: %s /delete-service → localhost:3003/delete-service", r.Method)
	models.Proxy = nil

	helpers.Proxy(golet_id, targetURL)
	models.Proxy.ServeHTTP(w, r)

}
