package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"frontend/helpers"
	"net/http"
	"time"

	"gorm.io/gorm"
)

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Println(payload)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		// http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		fmt.Println(err)
	}
}
func VerifyCodeBackendHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	code := r.FormValue("code")
	//  Decode JSON
	// Decide lookup field
	db := helpers.DbConn()
	var user User
	query := db.Model(&User{})
	query = query.Where("username = ?", username)
	result := query.First(&user)
	fmt.Println(result)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"success": true,
			"message": "Incorrect verification code",
		})
		return
	}
	//  Check code
	isCodeValid := user.VerifyCode == code
	isCodeNotExpired := user.VerifyCodeExpiry.After(time.Now())

	if isCodeValid && isCodeNotExpired {
		// Update user
		db.Model(&user).Update("is_verified", true)
		w.Header().Set("HX-Redirect", "/user")
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"success": true,
			"message": "Account verified successfully",
		})

		return
	}

	// Expired
	if !isCodeNotExpired {
		writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Verification code expired. Please sign up again.",
		})
		return
	}

	//  Incorrect code
	writeJSON(w, http.StatusBadRequest, map[string]interface{}{
		"success": false,
		"message": "Incorrect verification code",
	})
}
