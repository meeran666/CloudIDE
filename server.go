package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"project/components"
	"project/models"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	// Allow all origins for dev — restrict in production
	CheckOrigin: func(r *http.Request) bool { return true },
}

type OutputMessage struct {
	Dir_porfiles []models.Dirprofile `json:"dir_porfiles"`
	Error        bool                `json:"error"`
}
type InputMessage struct {
	Path string `json:"path"`
}

func sendOutput(conn *websocket.Conn, dir_porfiles []models.Dirprofile, isError bool) {
	msg := OutputMessage{Dir_porfiles: dir_porfiles, Error: isError}
	data, _ := json.Marshal(msg)
	fmt.Println(string(data))
	conn.WriteMessage(websocket.TextMessage, data)
}
func filelist(path string) ([]fs.DirEntry, error) {
	dir := "/home/zenith/meeran/cloudIDE/user1"
	subpath := dir + "/" + path
	root := os.DirFS(subpath)
	entries, err := fs.ReadDir(root, ".")
	if err != nil {
		log.Println("reading Directory error:", err)
		return nil, err
	}
	return entries, nil
}
func wsHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP → WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	log.Println("Client connected:", r.RemoteAddr)

	// Send welcome message
	for {
		// Read message from browser
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		var input InputMessage
		if err := json.Unmarshal(msg, &input); err != nil {
			fmt.Println("error")
			sendOutput(conn, []models.Dirprofile{}, true)
			continue
		}

		output, err := filelist("fold")
		for _, entry := range output {
			models.DirprofileArr = append(models.DirprofileArr, models.Dirprofile{Name: entry.Name(), IsDir: entry.IsDir()})
		}
		components.FileStructure(models.DirprofileArr).Render(r.Context(), w)
		models.DirprofileArr = nil
	}

	log.Println("Client disconnected")
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		models.DirprofileArr = nil
		if len(models.DirprofileArr) != 0 {
			fmt.Println("wo")
			component := components.Base(models.DirprofileArr)
			component.Render(r.Context(), w)
			return
		}
		dir := "/home/zenith/meeran/cloudIDE/user1"
		root := os.DirFS(dir)

		entries, err := fs.ReadDir(root, ".")

		if err != nil {
			log.Fatal(err)
		}

		for _, entry := range entries {
			models.DirprofileArr = append(models.DirprofileArr, models.Dirprofile{Name: entry.Name(), IsDir: entry.IsDir()})
		}
		components.Base(models.DirprofileArr).Render(r.Context(), w)
		models.DirprofileArr = nil

	})
	mux.HandleFunc("GET /ws", wsHandler)
	mux.Handle("GET /public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	fmt.Println("🚀  GOTTH Todo → http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
