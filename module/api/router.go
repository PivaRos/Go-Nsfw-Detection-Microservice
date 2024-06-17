package api

import (
	"net/http"

	"github.com/pivaros/go-image-recognition/utils"
)

func loadRoutes(router *http.ServeMux, AppState *utils.AppState) {

	router.HandleFunc("GET /as", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		return
	})
}
