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

	helpers.Proxy(golet_id, targetURL)

	models.Proxy.ServeHTTP(w, r)

	log.Println("Proxy running on :3000 → forwarding to :3001")
}
