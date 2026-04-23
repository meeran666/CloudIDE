package main

import (
	"bufio"
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

func applyLineChange(change LineChange, golet_id string) error {
	// read existing file

	//prod part
	// file_path := baseDir + "/" + golet_id + change.File_path

	//dev part
	file_path := baseDir + "/" + "user1" + change.File_path

	data, err := os.ReadFile(file_path)
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")

	// ensure slice is large enough
	if change.Line-1 >= len(lines) {
		for len(lines) <= change.Line-1 {
			lines = append(lines, "")
		}
	}

	// apply change (1-based index)
	lines[change.Line-1] = change.Content

	// write back
	output := strings.Join(lines, "\n")
	return os.WriteFile(file_path, []byte(output), 0644)
}
func SaveProgress(w http.ResponseWriter, r *http.Request) {
	golet_id := r.FormValue("golet_id")
	defer r.Body.Close()

	scanner := bufio.NewScanner(r.Body)

	for scanner.Scan() {
		line := scanner.Text()

		// remove trailing comma if present
		line = strings.TrimSuffix(line, ",")

		if strings.TrimSpace(line) == "" {
			continue
		}

		var change LineChange
		err := json.Unmarshal([]byte(line), &change)
		if err != nil {
			fmt.Println("Invalid JSON:", err)
			continue
		}

		err = applyLineChange(change, golet_id)
		if err != nil {
			fmt.Println("Apply error:", err)
		}
	}

	w.Write([]byte("Changes applied"))
}
