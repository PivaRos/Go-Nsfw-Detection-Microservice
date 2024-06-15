package api

import (
	"log"
	"net/http"

	"github.com/pivaros/go-image-recognition/api/middleware"
	"github.com/pivaros/go-image-recognition/utils"
)

func Run(appState *utils.AppState) error {
	//init the env variables
	Stack := middleware.CreateStack()
	mainRouter := http.NewServeMux()
	//create server instance
	loadRoutes(mainRouter, appState)
	app := http.Server{
		Addr:    ":" + appState.Env.PORT,
		Handler: Stack(mainRouter),
	}
	//start listening for traffic
	log.Println("Started Api on port " + appState.Env.PORT)
	err := app.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
