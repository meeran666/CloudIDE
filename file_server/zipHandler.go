package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func ZipHandler(w http.ResponseWriter, r *http.Request) {
	// Step 1: Read query param
	folderName := r.URL.Query().Get("folder_name")
	if folderName == "" {
		http.Error(w, "folder_name is required", http.StatusBadRequest)
		return
	}
	// Step 2: Build safe path
	folderPath := filepath.Join(baseDir, folderName)
	cleanPath := filepath.Clean(folderPath)

	// Security check (prevent ../ attack)
	if !strings.HasPrefix(cleanPath, baseDir) {
		http.Error(w, "invalid folder path", http.StatusBadRequest)
		return
	}

	// Check if folder exists
	info, err := os.Stat(cleanPath)
	if err != nil || !info.IsDir() {
		http.Error(w, "folder not found", http.StatusNotFound)
		return
	}

	// Step 3: Set response headers
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s.zip"`, folderName))

	// Step 4: Create zip writer (streaming)
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	// Step 5: Walk through folder and add files
	err = filepath.Walk(cleanPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Create relative path inside zip
		relPath, err := filepath.Rel(cleanPath, path)
		if err != nil {
			return err
		}

		// Skip root folder itself
		if relPath == "." {
			return nil
		}

		// If directory, just create header
		if info.IsDir() {
			_, err := zipWriter.Create(relPath + "/")
			return err
		}

		// Open file
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Create zip entry
		writer, err := zipWriter.Create(relPath)
		if err != nil {
			return err
		}

		// Copy file content into zip
		_, err = io.Copy(writer, file)
		return err
	})

	if err != nil {
		http.Error(w, "error creating zip: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
