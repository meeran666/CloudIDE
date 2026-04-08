package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"time"

	"frontend/helpers"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var db *gorm.DB

// ─────────────────────────────────────────────────────────────

type User struct {
	ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username         string    `gorm:"uniqueIndex"`
	Email            string    `gorm:"uniqueIndex"`
	Password         string
	VerifyCode       string
	VerifyCodeExpiry time.Time
	IsVerified       bool
	CreatedAt        time.Time
}

// ─────────────────────────────────────────────────────────────

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// ─────────────────────────────────────────────────────────────

func jsonResponse(w http.ResponseWriter, status int, payload RegisterResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func generateVerifyCode() string {
	return fmt.Sprintf("%06d", 100000+rand.Intn(900000))
}

// ─────────────────────────────────────────────────────────────

func sendVerificationEmail(email, username, code string) error {
	from := "you@example.com"
	password := "your-smtp-password"
	host := "smtp.example.com"
	port := "587"

	body := fmt.Sprintf(
		"Subject: Verify your account\n\nHi %s,\n\nYour verification code is: %s\n",
		username, code,
	)

	auth := smtp.PlainAuth("", from, password, host)
	return smtp.SendMail(host+":"+port, auth, from, []string{email}, []byte(body))
}

// ─────────────────────────────────────────────────────────────

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	db := helpers.DbConn()
	// Auto migrate (optional but useful)
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatal("Migration failed:", err)
	}
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonResponse(w, http.StatusBadRequest, RegisterResponse{false, "Invalid request body"})
		return
	}

	// ── 1. Check username ─────────────────────────────
	var existingUser User
	err := db.Where("username = ?", req.Username).First(&existingUser).Error
	if err == nil {
		if existingUser.IsVerified {
			jsonResponse(w, http.StatusBadRequest, RegisterResponse{false, "Username is already taken"})
			return
		}
	}
	// ── 2. Hash password ─────────────────────────────
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		log.Println("bcrypt error:", err)
		jsonResponse(w, http.StatusInternalServerError, RegisterResponse{false, "Error registering user"})
		return
	}
	hashedPassword := string(hashedBytes)

	verifyCode := generateVerifyCode()

	// ── 3. Check email ───────────────────────────────
	var userByEmail User
	err = db.Where("email = ?", req.Email).First(&userByEmail).Error

	if err == nil {
		// email exists
		if userByEmail.IsVerified {
			jsonResponse(w, http.StatusBadRequest, RegisterResponse{false, "User already exists with this email"})
			return
		}

		// update unverified user
		userByEmail.Username = req.Username
		userByEmail.Password = hashedPassword
		userByEmail.VerifyCode = verifyCode
		userByEmail.VerifyCodeExpiry = time.Now().Add(time.Hour)
		if err := db.Save(&userByEmail).Error; err != nil {
			log.Println("DB error (update):", err)
			jsonResponse(w, http.StatusInternalServerError, RegisterResponse{false, "Error registering user"})
			return
		}

	} else {
		// new user
		newUser := User{
			Username:         req.Username,
			Email:            req.Email,
			Password:         hashedPassword,
			VerifyCode:       verifyCode,
			VerifyCodeExpiry: time.Now().Add(10 * time.Hour),
			IsVerified:       false,
		}
		if err := db.Create(&newUser).Error; err != nil {
			log.Println("DB error (insert):", err)
			jsonResponse(w, http.StatusInternalServerError, RegisterResponse{false, "Error registering user"})
			return
		}

	}

	// ── 4. Send email ────────────────────────────────
	if err := sendVerificationEmail(req.Email, req.Username, verifyCode); err != nil {
		log.Println("Email error:", err)
		jsonResponse(w, http.StatusInternalServerError, RegisterResponse{false, "Failed to send verification email"})
		return
	}

	jsonResponse(w, http.StatusCreated, RegisterResponse{
		Success: true,
		Message: "User registered successfully. Please verify your account.",
	})

}
