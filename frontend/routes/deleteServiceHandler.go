package routes

import (
	"frontend/helpers"
	"frontend/models"
	"log"
	"net/http"
)

func ByPassToDeleteServiceHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Proxying: %s /save-change → localhost:3003/save-change", r.Method)
	models.Proxy = nil
	helpers.Proxy()
	models.Proxy.ServeHTTP(w, r)
}
