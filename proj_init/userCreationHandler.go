package main

import (
	"fmt"
	"net/http"
)

func UserCreationHandler(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	destination := "../user_environment/" + username // change this
	err := createDistination(destination)
	if err != nil {
		fmt.Println("Copy failed:", err)
		return
	}
}
