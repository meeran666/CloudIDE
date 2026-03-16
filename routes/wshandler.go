package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"project/models"
	"runtime"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	// Allow all origins for dev — restrict in production
	CheckOrigin: func(r *http.Request) bool { return true },
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
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
		var input models.InputMessage
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
	msg := models.OutputMessage{Output: text, Error: isError}
	data, _ := json.Marshal(msg)
	conn.WriteMessage(websocket.TextMessage, data)
}
