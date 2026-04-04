package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"project/models"
	"runtime"

	"github.com/creack/pty"
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

	// Launch a shell with PTY
	shell := "/bin/sh"
	if runtime.GOOS == "windows" {
		// pty does not support Windows; fall back or handle separately
		sendOutput(conn, "PTY not supported on Windows\r\n", true)
		return
	}
	if sh := os.Getenv("SHELL"); sh != "" {
		shell = sh
	}

	cmd := exec.Command(shell)
	cmd.Env = append(os.Environ(), "TERM=xterm-256color")

	// Start command in a PTY
	ptmx, err := pty.Start(cmd)
	if err != nil {
		log.Println("PTY start error:", err)
		sendOutput(conn, "Failed to start terminal: "+err.Error()+"\r\n", true)
		return
	}
	defer func() {
		ptmx.Close()
		cmd.Wait()
	}()

	// PTY → WebSocket: stream terminal output to browser
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := ptmx.Read(buf)
			fmt.Println(n)
			if n > 0 {
				msg := models.OutputMessage{
					Output: string(buf[:n]),
					Error:  false,
				}
				data, _ := json.Marshal(msg)
				if writeErr := conn.WriteMessage(websocket.TextMessage, data); writeErr != nil {
					log.Println("WebSocket write error:", writeErr)
					return
				}
			}
			if err != nil {
				if err != io.EOF {
					log.Println("PTY read error:", err)
				}
				return
			}
		}
	}()

	// WebSocket → PTY: forward browser keystrokes/commands to the shell
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}

		var input models.InputMessage
		if err := json.Unmarshal(msg, &input); err != nil {
			// If not JSON, treat raw bytes as terminal input (useful for xterm.js raw mode)
			if _, writeErr := ptmx.Write(msg); writeErr != nil {
				log.Println("PTY write error:", writeErr)
				break
			}
			continue
		}

		// Handle terminal resize
		if input.Cols > 0 && input.Rows > 0 {
			pty.Setsize(ptmx, &pty.Winsize{
				Cols: uint16(input.Cols),
				Rows: uint16(input.Rows),
			})
			continue
		}

		// Write command/input to PTY
		if input.Command != "" {
			if _, writeErr := ptmx.Write([]byte(input.Command)); writeErr != nil {
				log.Println("PTY write error:", writeErr)
				break
			}
		}
	}

	log.Println("Client disconnected")
}

func sendOutput(conn *websocket.Conn, text string, isError bool) {
	msg := models.OutputMessage{Output: text, Error: isError}
	data, _ := json.Marshal(msg)
	conn.WriteMessage(websocket.TextMessage, data)
}
