package routes

import (
	"frontend/components"
	"net/http"
)

func HomepageHandler(w http.ResponseWriter, r *http.Request) {
	components.HomepageBase().Render(r.Context(), w)

}
