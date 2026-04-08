package main

import (
	"fmt"
	"frontend/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Message from the browser
func Protected(w http.ResponseWriter, r *http.Request) {
	fmt.Println("value79")
	os.Exit(0)

	w.Write([]byte("You are authorized!"))
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/signup", routes.SignupHandler).Methods("POST")
	mux.HandleFunc("/signin", routes.SigninHandler).Methods("GET")

	mux.HandleFunc("/signinBackend", routes.SigninBackendHandler).Methods("POST")
	mux.HandleFunc("/dashboard", routes.AuthMiddleware(Protected)).Methods("GET")
	mux.HandleFunc("/", routes.LandingPageHandler).Methods("GET")

	mux.HandleFunc("/user", routes.AuthMiddleware(routes.UserHandler)).Methods("GET")
	mux.HandleFunc("/start", routes.IDEHandler).Methods("POST")
	mux.HandleFunc("/create", routes.IDEHandler).Methods("POST")
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
