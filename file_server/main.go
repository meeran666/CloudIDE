package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const baseDir = "../user_environment" // root directory on server

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/get-folder-zip", ZipHandler).Methods("GET")
	mux.HandleFunc("/save-progress", SaveProgress).Methods("POST")

	port := "3004"
	fmt.Println("🚀 Listening on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
