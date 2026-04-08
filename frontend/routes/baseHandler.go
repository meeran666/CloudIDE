package routes

import (
	"frontend/components"
	"net/http"
)

func BaseHandler(w http.ResponseWriter, r *http.Request) {

	components.Base(components.User()).Render(r.Context(), w)

}
