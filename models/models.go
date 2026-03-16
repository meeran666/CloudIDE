package models

type Dirprofile struct {
	Name  string
	IsDir bool
}
type FileRequest struct {
	Path string `json:"path"`
}

var DirprofileArr = []Dirprofile{}
