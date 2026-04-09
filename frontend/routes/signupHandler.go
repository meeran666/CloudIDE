package routes

import (
	"frontend/components"
	"net/http"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	components.Base(components.SignupPage()).Render(r.Context(), w)

}
