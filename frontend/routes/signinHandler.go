package routes

import (
	"frontend/components"
	"net/http"
)

func isHTMX(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}
func SigninHandler(w http.ResponseWriter, r *http.Request) {
	if isHTMX(r) {
		// Return ONLY signin component
		components.Signin().Render(r.Context(), w)
	} else {
		// Return full page (with navbar)
		components.Base(components.Signin()).Render(r.Context(), w)
		// components.Signin().Render(r.Context(), w)
	}

}
