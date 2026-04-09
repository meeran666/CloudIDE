package routes

import (
	"frontend/helpers"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("my_secret_key")

func SigninBackendHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}

	identifier := r.FormValue("identifier")
	password := r.FormValue("password")

	db := helpers.DbConn()

	// 2. Find user in DB
	var user User
	err = db.Where("username = ?", identifier).First(&user).Error
	if err != nil {

		// user not found OR DB error
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// 3. Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {

		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// 4. (Optional but recommended)
	if !user.IsVerified {
		http.Error(w, "Account not verified", http.StatusUnauthorized)
		return
	}
	// 5. Create JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		HttpOnly: true,
		Secure:   false, // true in production (HTTPS)
		Path:     "/",
		MaxAge:   86400, // 24 hours
	})

	w.WriteHeader(http.StatusOK)
}
