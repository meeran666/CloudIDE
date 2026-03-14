package main

import (
	"fmt"
	"log"
	"net/http"
	"project/components"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		component := components.Hello("water")
		component.Render(r.Context(), w)

	})
	mux.Handle("GET /public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	fmt.Println("🚀  GOTTH Todo → http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
