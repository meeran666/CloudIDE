package routes

import (
	"fmt"
	"frontend/helpers"
	"frontend/models"
	"net/http"
	"net/url"
	"strings"
)

func UserRecreationHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	workspace_name := r.FormValue("workspace_name")

	//extracting the golet_id of this workspace
	db := helpers.DbConn()
	var stack_info models.Stack
	err := db.Where(&models.Stack{WorkspaceName: workspace_name, Username: username}).First(&stack_info).Error
	if err != nil {
		// Found a verified user
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"success": false, "message": "Username is already taken"}`)

		return
	}
	data := url.Values{}

	golet_id := stack_info.Username + "-" + stack_info.GoletID.String()

	data.Set("stack", stack_info.Stack)
	data.Set("golet_id", golet_id)
	// 2. Target URL
	apiUrl := "http://localhost:3002/user_recreation"

	// 3. Create request with encoded body
	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println("Error:", err)
		http.Error(w, ":"+err.Error(), 400)
		return

	}
	// 4. Set mandatory header
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// 5. Send
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		http.Error(w, ":"+err.Error(), 400)
		return
	}

	w.Header().Set("HX-Redirect", "/start?golet_id="+golet_id+"&workspace_name="+stack_info.WorkspaceName)

	defer resp.Body.Close()

}
