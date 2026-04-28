package helpers

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Mu   sync.Mutex
}

var Clients = struct {
	sync.RWMutex
	M map[string]*Client
}{
	M: make(map[string]*Client),
}

// ✅ Add client
func AddClient(userID string, conn *websocket.Conn) {
	Clients.Lock()
	defer Clients.Unlock()

	Clients.M[userID] = &Client{Conn: conn}
}

// ✅ Get client
func GetClient(userID string) (*Client, bool) {
	Clients.RLock()
	defer Clients.RUnlock()

	c, ok := Clients.M[userID]
	return c, ok
}

// ✅ Remove client
func RemoveClient(userID string) {
	Clients.Lock()
	defer Clients.Unlock()

	delete(Clients.M, userID)
}
