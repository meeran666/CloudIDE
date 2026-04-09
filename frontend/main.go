package main

import (
	"fmt"
	"frontend/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/signupBackend", routes.SignupBackendHandler).Methods("POST")
	mux.HandleFunc("/signup", routes.SignupHandler).Methods("GET")
	mux.HandleFunc("/check-username-unique", routes.CheckUsernameUniqueHandler).Methods("GET")

	mux.HandleFunc("/signin", routes.SigninHandler).Methods("GET")
	mux.HandleFunc("/verifyAccount", routes.VerifyAccountHandler).Methods("GET")
	mux.HandleFunc("/verify-code-backend", routes.VerifyCodeBackendHandler).Methods("POST")
	mux.HandleFunc("/signinBackend", routes.SigninBackendHandler).Methods("POST")
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
