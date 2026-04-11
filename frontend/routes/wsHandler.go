package routes

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	// Allow all origins for dev — restrict in production
	CheckOrigin: func(r *http.Request) bool { return true },
}

func connectToServiceServer(golet_id string) (*websocket.Conn, error) {
	targetURL := "ws://" + golet_id + ".localhost:3006/ws"

	conn, _, err := websocket.DefaultDialer.Dial(targetURL, nil)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
func WsHandler(w http.ResponseWriter, r *http.Request) {

	golet_id := "service"
	// Upgrade frontend connection
	clientConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer clientConn.Close()

	// Connect to Server B
	serverConn, err := connectToServiceServer(golet_id)
	if err != nil {
		clientConn.WriteMessage(websocket.TextMessage, []byte("failed to connect backend"))
		return
	}
	defer serverConn.Close()

	// 🔁 Bidirectional pipe

	// Client → Server B
	go func() {
		for {
			_, msg, err := clientConn.ReadMessage()
			if err != nil {
				break
			}
			serverConn.WriteMessage(websocket.TextMessage, msg)
		}
	}()

	// Server B → Client
	for {
		_, msg, err := serverConn.ReadMessage()
		if err != nil {
			break
		}
		clientConn.WriteMessage(websocket.TextMessage, msg)
	}
}
