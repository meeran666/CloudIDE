package routes

import (
	"frontend/helpers"
	"frontend/models"
	"log"
	"net/http"
)

func FileDeleteProxy(w http.ResponseWriter, r *http.Request) {
	golet_id := r.FormValue("golet_id")
	//dev part
	// targetURL := "http://localhost:3003"
	//prod part
	targetURL := "http://" + golet_id + ".localhost:3006"

	log.Printf("Proxying: %s /file_delete → localhost:3006/file_delete", r.Method)
	models.Proxy = nil

	helpers.Proxy(golet_id, targetURL)

	models.Proxy.ServeHTTP(w, r)

	log.Println("Proxy running on :3000 → forwarding to :3001")
}
