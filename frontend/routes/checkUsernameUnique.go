package routes

import (
	"fmt"
	"frontend/helpers"
	"net/http"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Request query struct
type UsernameQuery struct {
	Username string `validate:"required,min=3,max=30,alphanum"`
}

// Validator instance
var validate = validator.New()

func CheckUsernameUniqueHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameter
	username := r.URL.Query().Get("username")
	query := UsernameQuery{Username: username}
	// Validate username
	err := validate.Struct(query)
	if err != nil {
		// Collect validation errors

		var errorMsg string
		for _, e := range err.(validator.ValidationErrors) {
			errorMsg += fmt.Sprintf("%s is invalid; ", e.Field())
		}
		if errorMsg == "" {
			errorMsg = "Invalid query parameters"
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"success": false, "message": "%s"}`, errorMsg)
		return
	}

	// Check if username exists and is verified
	db := helpers.DbConn()
	var user User
	err = db.Where("username = ? AND is_verified = ?", username, true).First(&user).Error
	if err == nil {
		// Found a verified user
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"success": false, "message": "Username is already taken"}`)
		return
	} else if err != gorm.ErrRecordNotFound {
		// Database error
		fmt.Println("Error checking username:", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"success": false, "message": "Error checking username"}`)
		return
	}

	// Username is unique

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"success": true, "message": "Username is unique"}`)
}
