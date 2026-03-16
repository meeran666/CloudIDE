package main

import (
	"fmt"
	"log"
	"net/http"
	"project/routes"
)

// Message from the browser

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", routes.HomepageHandler)
	mux.HandleFunc("/ws", routes.WsHandler)
	mux.HandleFunc("/browse", routes.BrowseHandler)
	mux.HandleFunc("/file", routes.FileHandler)
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	fmt.Println("🚀  GOTTH Todo → http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
