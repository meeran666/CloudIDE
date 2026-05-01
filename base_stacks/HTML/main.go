package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fs := http.FileServer(http.Dir("./public"))

	router := mux.NewRouter()

	// Serve all static files
	router.PathPrefix("/").Handler(fs)

	port := "3000"
	fmt.Println("🚀 Listening on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
