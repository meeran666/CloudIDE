package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runner/routes"
	"time"

	"github.com/gorilla/mux"
)

// Message from the browser
func notify() {
	golet_id := os.Getenv("GOLET_ID")
	workspace_name := os.Getenv("WORKSPACE_NAME")

	if golet_id == "" || workspace_name == "" {
		fmt.Println("golet_id or workspace_name is not set")
		return
	}
	fmt.Println("golet_id", golet_id)
	remoteServerURL := "http://100.66.155.68:3001/notify" + "?golet_id=" + golet_id + "&workspace_name=" + workspace_name
	resp, err := http.Get(remoteServerURL)
	fmt.Println("err", err)
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	defer resp.Body.Close()
}
func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", routes.HomepageHandler).Methods("GET")
	mux.HandleFunc("/ws", routes.WsHandler).Methods("GET")
	mux.HandleFunc("/browse", routes.BrowseHandler).Methods("GET")
	mux.HandleFunc("/file", routes.FileReadHandler).Methods("POST")
	mux.HandleFunc("/save_change", routes.ChangeHandler).Methods("POST")
	mux.HandleFunc("/delete_service", routes.DeleteServiceHandler).Methods("POST")
	mux.HandleFunc("/file_create", routes.FileCreateHandler).Methods("POST")
	mux.HandleFunc("/file_delete", routes.FileDeleteHandler).Methods("POST")

	port := "3006"
	fmt.Println("🚀 Listening on port:", port)

	//  Run notify in background
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Calling notify...")
		notify()
	}()

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
