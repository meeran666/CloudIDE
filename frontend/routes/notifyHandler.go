package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"frontend/components"
	"frontend/helpers"
	"frontend/models"
	"net/http"

	"github.com/gorilla/websocket"
)

func NotifyHandler(w http.ResponseWriter, r *http.Request) {
	golet_id := r.FormValue("golet_id")
	workspace_name := r.URL.Query().Get("workspace_name")

	// prod part
	targetURL := "http://" + golet_id + ".localhost:3006"
	// dev part
	// targetURL := "http://localhost:3003"
	resp, err := http.Get(targetURL)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	models.DirprofileArr = nil
	err = json.NewDecoder(resp.Body).Decode(&models.DirprofileArr)
	if err != nil {
		fmt.Println(err)
	}

	client, ok := helpers.GetClient(golet_id)
	if !ok {
		http.Error(w, "No active websocket", 400)
		return
	}

	var buf bytes.Buffer
	components.Base(components.IDEBase(models.DirprofileArr, golet_id, workspace_name)).Render(r.Context(), &buf)

	// err := components.Notification().Render(r.Context(), &buf)
	if err != nil {
		http.Error(w, "render error", 500)
		return
	}

	client.Mu.Lock()
	defer client.Mu.Unlock()

	err = client.Conn.WriteMessage(websocket.TextMessage, buf.Bytes())
	if err != nil {
		http.Error(w, "Write failed", 500)
		return
	}

	w.Write([]byte("sent"))
}
