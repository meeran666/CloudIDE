package routes

import (
	"frontend/components"
	"net/http"
)

func LandingPageHandler(w http.ResponseWriter, r *http.Request) {
	components.Base(components.LandingPage()).Render(r.Context(), w)

}
