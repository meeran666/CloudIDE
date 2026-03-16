package models

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
}

// Message sent back to browser
type OutputMessage struct {
	Output string `json:"output"`
	Error  bool   `json:"error"`
}
