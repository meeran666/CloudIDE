package main

import (
	"fmt"
	"log"
	"net/http"
)

// Message from the browser

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomepageHandler)
	fmt.Println("🚀  GOTTH Todo → http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
