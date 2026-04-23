package routes

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"path/filepath"
)

type CreateRequest struct {
	Path string `json:"path"`
	Name string `json:"name"`
	Dir  bool   `json:"dir"`
}

func createFile(path string) error {
	// Check if already exists
	if _, err := os.Stat(path); err == nil {
		return errors.New("file already exists")
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}
func FileCreateHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Prevent path traversal (important)

	path := filepath.Clean(req.Path)
	fullPath := filepath.Join("/workspace", path, req.Name)
	if req.Dir {
		err = os.MkdirAll(fullPath, 0755)
	} else {
		err = createFile(fullPath)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("created successfully"))
}
