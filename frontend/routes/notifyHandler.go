package routes

import (
	"encoding/json"
	"fmt"
	"frontend/helpers"
	"frontend/models"
	"net/http"

	"github.com/gorilla/websocket"
)

func NotifyHandler(w http.ResponseWriter, r *http.Request) {
	golet_id := r.FormValue("golet_id")
	workspace_name := r.FormValue("workspace_name")
	targetURL := "http://" + golet_id + ".localhost:3006"

	resp, err := http.Get(targetURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	models.DirprofileArr = nil
	err = json.NewDecoder(resp.Body).Decode(&models.DirprofileArr)
	if err != nil {
		fmt.Println(err)
		return
	}

	client, ok := helpers.GetClient(golet_id)
	if !ok {
		http.Error(w, "No active websocket", 400)
		return
	}
	// Create payload instead of rendering HTML
	payload := map[string]interface{}{
		"golet_id":       golet_id,
		"workspace_name": workspace_name,
	}

	data, err := json.Marshal(payload)
	fmt.Println("payload", payload)
	if err != nil {
		http.Error(w, "json error", 500)
		return
	}

	client.Mu.Lock()
	defer client.Mu.Unlock()
	err = client.Conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		http.Error(w, "Write failed", 500)
		return
	}

	w.Write([]byte("sent"))
}
