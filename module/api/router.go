package api

import (
	"net/http"

	"github.com/pivaros/go-image-recognition/utils"
)

func loadRoutes(router *http.ServeMux, AppState *utils.AppState) {

	// //auth
	// routes.RegisterAuthRoutes(router, AppState)
	// //admin
	// routes.RegisterAdminRoutes(router, AppState)
	// //appointment
	// routes.RegisterAppointmentRoutes(router, AppState)
}
