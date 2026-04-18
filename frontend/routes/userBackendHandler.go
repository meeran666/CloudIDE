package routes

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func UserBackendHandler(w http.ResponseWriter, r *http.Request) {
	// sending the request to orchestrator server for creating user space folder
	data := url.Values{}
	golet_id := r.FormValue("golet_id")
	stack := r.FormValue("stack")

	data.Set("golet_id", golet_id)
	data.Set("stack", stack)
	// 2. Target URL
	apiUrl := "http://localhost:3003/user_creation"

	// 3. Create request with encoded body
	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println("Error:", err)
		http.Error(w, ":"+err.Error(), 400)

	}
	// 4. Set mandatory header
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// 5. Send
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		http.Error(w, ":"+err.Error(), 400)

	}

	defer resp.Body.Close()
	w.Header().Set("HX-Redirect", "/start?golet_id="+golet_id)

	fmt.Println("Status:", resp.Status)
}
