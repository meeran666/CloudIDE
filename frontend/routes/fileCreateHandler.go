package routes

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"path/filepath"
)

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

	err := json.NewDecoder(r.Body).Decode(&req)
	path := r.FormValue("path")
	name := r.FormValue("name ")
	dir := r.FormValue("dir")
	if err != nil {

		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Basic validation
	if req.Path == "" || req.Name == "" {
		http.Error(w, "path and name are required", http.StatusBadRequest)
		return
	}

	// Prevent path traversal (important)
	cleanBase := filepath.Clean(req.Path)
	fullPath := filepath.Join(cleanBase, req.Name)

	// Optional: restrict to workspace root
	// root := "./workspace"
	// if !strings.HasPrefix(fullPath, root) {
	//     http.Error(w, "invalid path", http.StatusForbidden)
	//     return
	// }

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
