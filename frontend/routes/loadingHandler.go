package routes

import (
	"frontend/components"
	"frontend/models"
	"net/http"
)

func LoadingHandler(w http.ResponseWriter, r *http.Request) {
	golet_id := r.URL.Query().Get("golet_id")
	workspace_name := r.URL.Query().Get("workspace_name")

	components.Base(components.Loading(golet_id, workspace_name)).Render(r.Context(), w)
	models.DirprofileArr = nil

}
