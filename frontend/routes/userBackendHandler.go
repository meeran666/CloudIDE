package routes

import (
	"errors"
	"fmt"
	"frontend/helpers"
	"frontend/models"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5/pgconn"
)

func updateStackInfo(workspace_name, stack, username string) (string, error) {
	db := helpers.DbConn()
	// Auto migrate
	// if err := db.AutoMigrate(&models.Stack{}); err != nil {
	// 	log.Fatal("Migration failed:", err)
	// }

	newStack := models.Stack{
		Username:      username,
		WorkspaceName: workspace_name,
		Stack:         stack,
		Lastupdated:   time.Now(),
	}
	if err := db.Create(&newStack).Error; err != nil {
		return "", err
	}

	return newStack.GoletID.String(), nil
}

func serviceInitHandler(golet_id, workspace_name, stack string) error {

	data := url.Values{}
	data.Set("golet_id", golet_id)
	data.Set("stack", stack)
	data.Set("workspace_name", workspace_name)
	// 2. Target URL
	apiUrl := "http://localhost:3002/user_creation"
	// 3. Create request with encoded body
	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return err

	}
	// 4. Set mandatory header
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// 5. Send
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {

		return err
	}

	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
	return nil
}

func UserBackendHandler(w http.ResponseWriter, r *http.Request) {
	workspace_name := r.FormValue("workspace_name")
	stack := r.FormValue("stack")
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	claims, ok := r.Context().Value(models.UserContextKey).(jwt.MapClaims)

	if !ok {
		http.Error(w, "No user data", http.StatusUnauthorized)
	}

	username := claims["username"].(string)
	// sending the request to orchestrator server for creating user space folder
	golet_id, err := updateStackInfo(workspace_name, stack, username)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.ConstraintName {
			case "idx_user_workspace":
				fmt.Println("workspace already exists for this name")
				http.Error(w, ":"+err.Error(), 409)

			}
		}
	}

	golet_id = username + "-" + golet_id
	// golet_id := "meeran-e20fce22-dc22-439f-901f-4bf59a71662d"

	helpers.Clients.Lock()
	helpers.Clients.M[golet_id] = &helpers.Client{Conn: conn}
	helpers.Clients.Unlock()

	err = serviceInitHandler(golet_id, workspace_name, stack)
	if err != nil {
		fmt.Println("Error:", err)
		http.Error(w, ":"+err.Error(), 400)
		return
	}

}
