package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/exec"
	"project/components"
	"project/models"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/gorilla/websocket"
)

// Message from the browser
type InputMessage struct {
	Command string `json:"command"`
}

// Message sent back to browser
type OutputMessage struct {
	Output string `json:"output"`
	Error  bool   `json:"error"`
}

var upgrader = websocket.Upgrader{
	// Allow all origins for dev — restrict in production
	CheckOrigin: func(r *http.Request) bool { return true },
}

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	models.DirprofileArr = nil
	err := filelist("user1")
	if err != nil {
		color.Red("Error: %v", err)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	components.Base(models.DirprofileArr).Render(r.Context(), w)
	models.DirprofileArr = nil

}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside")
	// Upgrade HTTP → WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	log.Println("Client connected:", r.RemoteAddr)

	// Send welcome message
	sendOutput(conn, "Connected to Go terminal. Type a command and press Enter.\r\n", false)

	for {
		// Read message from browser
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		fmt.Println(string(msg))
		var input InputMessage
		if err := json.Unmarshal(msg, &input); err != nil {
			sendOutput(conn, "Invalid message format\r\n", true)
			continue
		}

		cmd := strings.TrimSpace(input.Command)
		if cmd == "" {
			continue
		}

		log.Printf("Running command: %q", cmd)

		// Echo the command back (like a real terminal)
		// sendOutput(conn, fmt.Sprintf("$ %s\r\n", cmd), false)

		// Handle built-in "clear"
		if cmd == "clear" {
			sendOutput(conn, "\033[2J\033[H", false)
			continue
		}

		// Run command in shell
		output, isErr := runCommand(cmd)
		sendOutput(conn, output, isErr)
	}

	log.Println("Client disconnected")
}

func runCommand(cmd string) (string, bool) {
	var c *exec.Cmd

	if runtime.GOOS == "windows" {
		c = exec.Command("cmd", "/C", cmd)
	} else {
		c = exec.Command("sh", "-c", cmd)
	}

	// Combine stdout + stderr
	out, err := c.CombinedOutput()
	output := string(out)

	// Normalize line endings for xterm.js
	output = strings.ReplaceAll(output, "\n", "\r\n")
	if !strings.HasSuffix(output, "\r\n") {
		output += "\r\n"
	}

	if err != nil {
		// Exit error: show output + error message
		if output == "\r\n" {
			output = err.Error() + "\r\n"
		}
		return output, true
	}

	return output, false
}

func sendOutput(conn *websocket.Conn, text string, isError bool) {
	msg := OutputMessage{Output: text, Error: isError}
	data, _ := json.Marshal(msg)
	conn.WriteMessage(websocket.TextMessage, data)
}

func browseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	path := r.FormValue("path")
	err := filelist(path)
	if err != nil {
		color.Red("Error: %v", err)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	components.FileStructure(path, models.DirprofileArr).Render(r.Context(), w)
	models.DirprofileArr = nil
}
func fileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req models.FileRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		color.Red("Error: %v", err)
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	data, err := os.ReadFile(req.Path)
	fmt.Println(string(data))
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write(data)
}
func filelist(path string) error {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	subpath := dir + "/" + path
	root := os.DirFS(subpath)
	entries, err := fs.ReadDir(root, ".")
	if err != nil {
		return err
	}
	models.DirprofileArr = nil
	for _, entry := range entries {
		models.DirprofileArr = append(models.DirprofileArr, models.Dirprofile{Name: entry.Name(), IsDir: entry.IsDir()})
	}
	return nil
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homepageHandler)
	mux.HandleFunc("/ws", wsHandler)
	mux.HandleFunc("/browse", browseHandler)
	mux.HandleFunc("/file", fileHandler)
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	fmt.Println("🚀  GOTTH Todo → http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
