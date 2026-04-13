package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Message from the browser

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/user_creation", InitHandler).Methods("POST")
	port := "3003"
	fmt.Println("🚀 Listening on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
