package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type LineChange struct {
	Line      int    `json:"line"`
	Content   string `json:"content"`
	File_path string `json:"file_path"`
}

type ChangePayload struct {
	Lines []LineChange `json:"lines"`
}

func ChangeHandler(w http.ResponseWriter, r *http.Request) {
	var payload ChangePayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	path := "../user_environment/user1"

	// filePath := "../user_environment/user1/down.html"
	// filePath := path + payload.Lines[]
	// content, err := os.ReadFile(filePath)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// Apply each line change
	for _, change := range payload.Lines {

		filePath := path + change.File_path
		content, err := os.ReadFile(filePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		lines := strings.Split(string(content), "\n")

		lineIndex := change.Line - 1 // Convert to 0-based index
		if lineIndex < 0 {
			continue
		}

		// Extend slice if new lines are added beyond current length
		for len(lines) <= lineIndex {
			lines = append(lines, "")
		}

		lines[lineIndex] = change.Content
		updatedContent := strings.Join(lines, "\n")
		if err := os.WriteFile(filePath, []byte(updatedContent), 0644); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	fmt.Println("File updated successfully")
	w.WriteHeader(http.StatusOK)
}
