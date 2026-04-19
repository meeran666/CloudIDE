package routes

import (
	"frontend/helpers"
	"frontend/models"
	"log"
	"net/http"
)

func FileCreateProxy(w http.ResponseWriter, r *http.Request) {
	golet_id := r.FormValue("golet_id")
	//dev part
	targetURL := "http://localhost:3003"
	//prod part
	// targetURL := "http://" + golet_id + ".localhost:3000"

	log.Printf("Proxying: %s /file-create → localhost:3003/file-create", r.Method)
	models.Proxy = nil

	helpers.Proxy(golet_id, targetURL)

	models.Proxy.ServeHTTP(w, r)

	log.Println("Proxy running on :3000 → forwarding to :3001")
}
