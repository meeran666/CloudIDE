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
	mux.HandleFunc("/check_username_unique", routes.CheckUsernameUniqueHandler).Methods("GET")
	mux.HandleFunc("/signin", routes.SigninHandler).Methods("GET")
	mux.HandleFunc("/save_change", routes.ChangeHandler).Methods("POST")
	mux.HandleFunc("/verifyAccount", routes.VerifyAccountHandler).Methods("GET")
	mux.HandleFunc("/verify_code_backend", routes.VerifyCodeBackendHandler).Methods("POST")
	mux.HandleFunc("/signinBackend", routes.SigninBackendHandler).Methods("POST")
	mux.HandleFunc("/", routes.LandingPageHandler).Methods("GET")
	mux.HandleFunc("/user", routes.AuthMiddleware(routes.UserHandler)).Methods("GET")
	mux.HandleFunc("/user_recreation", routes.AuthMiddleware(routes.UserRecreationHandler)).Methods("POST")
	mux.HandleFunc("/file", routes.AuthMiddleware(routes.FileHandler)).Methods("POST")
	mux.HandleFunc("/ws", routes.AuthMiddleware(routes.WsHandler)).Methods("GET")
	mux.HandleFunc("/userBackend", routes.AuthMiddleware(routes.UserBackendHandler)).Methods("GET")
	mux.HandleFunc("/delete_service", routes.DeleteServiceProxy).Methods("POST")
	mux.PathPrefix("/preview").HandlerFunc(routes.PreviewProxy)
	mux.HandleFunc("/file_create", routes.FileCreateProxy).Methods("POST")
	mux.HandleFunc("/file_delete", routes.FileDeleteProxy).Methods("POST")

	mux.HandleFunc("/start", routes.AuthMiddleware(routes.IDEHandler)).Methods("GET")
	mux.HandleFunc("/notify", routes.NotifyHandler).Methods("GET")
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
