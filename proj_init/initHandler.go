package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type RequestBody struct {
	Golet_id string `json:"golet_id"`
	Stack    string `json:"stack"`
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

	// err := r.ParseForm()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	http.Error(w, err.Error(), 400)
	// 	return
	// }

	// // Get values
	// golet_id := r.FormValue("golet_id")
	// stack := r.FormValue("stack")

	// // golet_id := "my_computer_contain_golet"

	// source := "../base_stacks/" + stack
	// destination := "../user_environment/" + golet_id
	// err = createDistination(destination)
	// err = copyDirContents(source, destination)
	// os.Exit(0)
	// if err != nil {
	// 	fmt.Println("Copy failed:" + err.Error())
	// 	http.Error(w, "Copy failed:"+err.Error(), 400)

	// 	return
	// }
	// fmt.Println("All files and folders copied successfully")
	// err = ContainerHandler(golet_id)
	// if err != nil {
	// 	fmt.Println("service creation failed:" + err.Error())
	// 	http.Error(w, "service creation failed:"+err.Error(), 400)
	// 	return
	// }
	fmt.Println("service created successfully")
	w.WriteHeader(http.StatusCreated)
}
