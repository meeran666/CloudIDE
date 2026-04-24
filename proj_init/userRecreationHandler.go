package main

import (
	"fmt"
	"net/http"
)

func UserRecreationHandler(w http.ResponseWriter, r *http.Request) {
	golet_id := r.FormValue("golet_id")
	stack := r.FormValue("stack")
	err := ContainerCreater(golet_id, stack)
	if err != nil {
		fmt.Println("container creation failed:" + err.Error())
		http.Error(w, "container creation failed:"+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
