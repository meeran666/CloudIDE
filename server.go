package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"project/components"
	"project/models"

	"github.com/fatih/color"
)

func filelist(path string) error {

	dir := "/home/zenith/meeran/cloudIDE"
	subpath := dir + "/" + path
	root := os.DirFS(subpath)
	entries, err := fs.ReadDir(root, ".")
	if err != nil {
		return err
	}
	models.DirprofileArr = nil
	for _, entry := range entries {
		models.DirprofileArr = append(models.DirprofileArr, models.Dirprofile{Name: entry.Name(), IsDir: entry.IsDir()})
	}
	return nil
}
func browseHandler(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")
	err := filelist(path)
	if err != nil {
		color.Red("Error: %v", err)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	components.FileStructure(path, models.DirprofileArr).Render(r.Context(), w)
	models.DirprofileArr = nil
}
func homepageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	models.DirprofileArr = nil
	err := filelist("user1")
	if err != nil {
		color.Red("Error: %v", err)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	components.Base(models.DirprofileArr).Render(r.Context(), w)
	models.DirprofileArr = nil

}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	var req models.FileRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		color.Red("Error: %v", err)
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	data, err := os.ReadFile(req.Path)
	fmt.Println(string(data))
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write(data)
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", homepageHandler)
	mux.HandleFunc("POST /browse", browseHandler)
	mux.HandleFunc("POST /file", fileHandler)
	mux.Handle("GET /public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	fmt.Println("🚀  GOTTH Todo → http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
