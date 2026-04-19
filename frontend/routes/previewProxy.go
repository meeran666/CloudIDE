package routes

import (
	"frontend/helpers"
	"frontend/models"
	"log"
	"net/http"
)

func PreviewProxy(w http.ResponseWriter, r *http.Request) {
	golet_id := r.FormValue("golet_id")
	//dev part
	targetURL := "http://localhost:5501"
	//prod part
	// targetURL := "http://" + golet_id + ".localhost:3000"

	log.Printf("Proxying: %s /preview.html → localhost:5501/preview.html", r.Method)
	models.Proxy = nil

	helpers.Proxy(golet_id, targetURL)

	models.Proxy.ServeHTTP(w, r)

}
