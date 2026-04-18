package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

// const remoteServerURL = "http://192.168.1.5:3004/get-folder-zip"
const remoteServerURL = "http://localhost:3004/get-folder-zip"

func main() {
	// Step 1: Read env variable
	folderName := os.Getenv("BASE_FOLDER")
	if folderName == "" {
		fmt.Println("BASE_FOLDER not set")
		return
	}

	fmt.Println("Requesting folder:", folderName)

	// Step 2: Build request URL
	reqURL := fmt.Sprintf("%s?folder_name=%s",
		remoteServerURL,
		url.QueryEscape(folderName),
	)

	// Step 3: Send request
	resp, err := http.Get(reqURL)

	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Bad response:", resp.Status)
		return
	}

	// Step 4: Save ZIP file
	zipPath := "./folder.zip"
	outFile, err := os.Create(zipPath)
	if err != nil {
		fmt.Println("Failed to create zip file:", err)
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)

	if err != nil {
		fmt.Println("Failed to save zip:", err)
		return
	}

	fmt.Println("ZIP downloaded at:", zipPath)

	// Step 5: Extract ZIP
	//dev part
	destDir := "./workspace"
	//prod part
	// destDir := "/workspace"

	err = unzip(zipPath, destDir)
	if err != nil {
		fmt.Println("Unzip failed:", err)
		return
	}

	fmt.Println("Extraction completed to:", destDir)
	os.Exit(0)

	// Optional: Start your server after setup
}
