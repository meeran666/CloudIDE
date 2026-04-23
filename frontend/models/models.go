package models

import (
	"net/http/httputil"
	"time"

	"github.com/google/uuid"
)

type Dirprofile struct {
	Name  string
	IsDir bool
}
type FileRequest struct {
	Path string `json:"path"`
}

var DirprofileArr = []Dirprofile{}

type InputMessage struct {
	Command string `json:"command"`
	Rows    uint16 `json:"rows"`
	Cols    uint16 `json:"cols"`
}

// Message sent back to browser
type OutputMessage struct {
	Output string `json:"output"`
	Error  bool   `json:"error"`
}

type Stack struct {
	Username      string `gorm:"not null;uniqueIndex:idx_user_workspace"`
	WorkspaceName string `gorm:"not null;uniqueIndex:idx_user_workspace"`
	Stack         string `gorm:"check:stack IN ('React','Node');not null"`
	Lastupdated   time.Time
	GoletID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
}

var Proxy *httputil.ReverseProxy

type contextKey string

const UserContextKey = contextKey("user")
