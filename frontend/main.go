package main

import (
	"fmt"
	"frontend/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Message from the browser

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", routes.IDEHandler).Methods("GET")
	mux.HandleFunc("/browse", routes.BrowseHandler).Methods("POST")
	mux.PathPrefix("/public/").Handler(
		http.StripPrefix("/public/", http.FileServer(http.Dir("public"))),
	).Methods("GET")
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/favicon.ico.png")
	}).Methods("GET")

	port := "3001"
	fmt.Println("🚀 Listening on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
