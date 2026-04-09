package routes

import (
	"frontend/components"
	"net/http"
)

func VerifyAccountHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	components.Base(components.VerifyAccount(username)).Render(r.Context(), w)

}
