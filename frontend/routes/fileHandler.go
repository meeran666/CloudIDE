package routes

import (
	"bytes"
	"io"
	"net/http"
)

func FileHandler(w http.ResponseWriter, r *http.Request) {
	// Read incoming body
	golet_id := r.FormValue("golet_id")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read request", 400)
		return
	}
	defer r.Body.Close()

	// prod part
	// golet_id := "weber"
	targetURL := "http://" + golet_id + ".localhost:3006/file"

	//dev part
	// targetURL := "http://localhost:3003/file"

	req, err := http.NewRequest("POST", targetURL, bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, "failed to create request", 500)
		return
	}

	// Copy headers (important for JSON)
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "failed to reach target server", 500)
		return
	}
	defer resp.Body.Close()

	// Read response from target server
	respBody, _ := io.ReadAll(resp.Body)

	// Send response back to original client
	w.WriteHeader(resp.StatusCode)
	w.Write(respBody)
}
