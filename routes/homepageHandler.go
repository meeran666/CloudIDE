package routes

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"project/components"
	"project/models"

	"github.com/fatih/color"
)

func HomepageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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
func filelist(path string) error {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
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
