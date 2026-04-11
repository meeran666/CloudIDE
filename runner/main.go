package main

import (
	"fmt"
	"log"
	"net/http"
	"runner/routes"

	"github.com/gorilla/mux"
)

// Message from the browser

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", routes.HomepageHandler).Methods("GET")
	mux.HandleFunc("/ws", routes.WsHandler).Methods("GET")
	mux.HandleFunc("/browse", routes.BrowseHandler).Methods("GET")
	mux.HandleFunc("/file", routes.FileReadHandler).Methods("POST")
	port := "3003"
	fmt.Println("🚀 Listening on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
