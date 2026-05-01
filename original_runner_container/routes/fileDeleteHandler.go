package routes

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"runner/models"
	"strings"
)

type DeleteRequest struct {
	Path string `json:"path"`
}

func FileDeleteHandler(w http.ResponseWriter, r *http.Request) {
	var req DeleteRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if req.Path == "" {
		http.Error(w, "path is required", http.StatusBadRequest)
		return
	}

	baseDir := filepath.Clean(models.BaseDir)

	// Build full path safely
	fullPath := filepath.Join(baseDir, req.Path)

	// Security check
	rel, err := filepath.Rel(baseDir, fullPath)
	if err != nil || strings.HasPrefix(rel, "..") {
		http.Error(w, "invalid path", http.StatusForbidden)
		return
	}

	// Check if exists
	_, err = os.Stat(fullPath)
	if os.IsNotExist(err) {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	// Delete file or folder (recursive)
	err = os.RemoveAll(fullPath)
	if err != nil {
		http.Error(w, "failed to delete: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("deleted successfully"))
}
