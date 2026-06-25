package handlers

import (
	"net/http"

	"github.com/PlinyTheYounger0/pliny_personal_portfolio/static/templates"
)

func (cfg *ApiConfig) Index(w http.ResponseWriter, r *http.Request) {

	cfg.Logger.InfoContext(r.Context(), "Served: "+r.URL.Path)
	component := templates.Index("Home Page")

	err := component.Render(r.Context(), w)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to Load Index.", err)
	}
}
