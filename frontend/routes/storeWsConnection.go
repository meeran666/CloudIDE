package routes

import (
	"fmt"
	"frontend/helpers"
	"net/http"

	"github.com/gorilla/websocket"
)

func StoreWsConnection(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	golet_id := r.URL.Query().Get("golet_id")
	fmt.Println(golet_id)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	helpers.Clients.Lock()
	helpers.Clients.M[golet_id] = &helpers.Client{Conn: conn}
	helpers.Clients.Unlock()
}
