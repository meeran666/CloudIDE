package routes

import (
	"fmt"
	"frontend/components"
	"frontend/helpers"
	"frontend/models"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func stacksInfo(username string) error {
	var stacks []Stack
	db := helpers.DbConn()

	err := db.Where("username = ?", username).Find(&stacks).Error
	if err != nil {
		return err
	}

	return nil
}

func UserHandler(w http.ResponseWriter, r *http.Request) {

	claims, ok := r.Context().Value(models.UserContextKey).(jwt.MapClaims)

	if !ok {
		http.Error(w, "No user data", http.StatusUnauthorized)
	}

	username := claims["username"].(string)
	err := stacksInfo(username)
	if err != nil {
		fmt.Println("Error:", err)
		http.Error(w, ":"+err.Error(), 400)

	}
	components.Base(components.User()).Render(r.Context(), w)
}
