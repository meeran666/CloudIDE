package routes

import (
	"encoding/json"
	"io/fs"
	"net/http"
	"os"
	"runner/models"

	"github.com/fatih/color"
)

func filelist(path string) error {

	// dev part
	// dir, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// subpath := dir + "/" + models.BaseDir + "/" + path

	// prod part
	subpath := path

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

func HomepageHandler(w http.ResponseWriter, r *http.Request) {

	// dev part
	// golet_id := "weber"

	err := filelist("/workspace")
	if err != nil {
		color.Red("Error: %v", err)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	// Set response type
	w.Header().Set("Content-Type", "application/json")

	// Encode and return JSON
	json.NewEncoder(w).Encode(models.DirprofileArr)
	models.DirprofileArr = nil

}
