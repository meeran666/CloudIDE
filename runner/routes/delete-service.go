package routes

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DeleteServiceHandler(w http.ResponseWriter, r *http.Request) {
	golet_id := r.FormValue("golet_id")
	file, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//dev part
	// URL := "http://localhost:3004/save-progress?golet_id=" + golet_id
	//prod part
	URL := "http://100.66.155.68:3004/save-progress?golet_id=" + golet_id
	req, err := http.NewRequest("POST", URL, file)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "text/plain")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	err = os.Truncate("file.txt", 0)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Response:", string(body))
}
