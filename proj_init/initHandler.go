package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type RequestBody struct {
	User_id  string `json:"user_id"`
	Golet_id string `json:"golet_id"`
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	info, err := srcFile.Stat()
	if err != nil {
		return err
	}

	// create destination file
	dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, info.Mode())
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

func copyDirContents(srcDir, dstDir string) error {
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return err
	}

	// create destination folder if not exists
	err = os.MkdirAll(dstDir, 0755)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(srcDir, entry.Name())
		dstPath := filepath.Join(dstDir, entry.Name())

		if entry.IsDir() {
			err = copyDirContents(srcPath, dstPath)
			if err != nil {
				return err
			}
		} else {
			err = copyFile(srcPath, dstPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
func createDistination(path string) error {
	// path := "../user_environment" + golet_id
	fmt.Println(path)
	err := os.Mkdir(path, 0755)
	return err
}

func InitHandler(w http.ResponseWriter, r *http.Request) {
	var body RequestBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	golet_id := body.Golet_id
	// golet_id := "my_computer_contain_golet"
	source := "../base_stacks/node_stack" // folder whose contents you want to copy

	destination := "../user_environment/" + body.User_id // change this
	err = createDistination(destination)
	err = copyDirContents(source, destination)
	if err != nil {
		fmt.Println("Copy failed:", err)
		return
	}
	fmt.Println("All files and folders copied successfully")
	ContainerHandler(golet_id)

}
