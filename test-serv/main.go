package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	targetURL := "http://weber.localhost:3006/file_create"
	// targetURL := "http://weber.localhost:3006/browse?path=/mold"

	// Create payload
	payload := map[string]interface{}{
		"path": "/mold",
		"name": "rino.txt",
		"dir":  false,
	}

	// Convert to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	// Create request
	req, err := http.NewRequest("POST", targetURL, bytes.NewBuffer(jsonData))
	// req, err := http.NewRequest("POST", targetURL, bytes.NewBuffer(jsonData))

	if err != nil {
		panic(err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
}
