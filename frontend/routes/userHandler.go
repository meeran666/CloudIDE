package routes

import (
	"fmt"
	"frontend/components"
	"frontend/helpers"
	"frontend/models"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func stacksInfo(username string, stacks *[]models.Stack) error {
	db := helpers.DbConn()

	err := db.Where("username = ?", username).Find(&stacks).Error
	if err != nil {
		return err
	}

	return nil
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	var stacks []models.Stack

	claims, ok := r.Context().Value(models.UserContextKey).(jwt.MapClaims)

	if !ok {
		http.Error(w, "No user data", http.StatusUnauthorized)
	}

	username := claims["username"].(string)
	err := stacksInfo(username, &stacks)
	if err != nil {
		fmt.Println("Error:", err)
		http.Error(w, ":"+err.Error(), 400)

	}
	components.Base(components.User(stacks)).Render(r.Context(), w)
}
