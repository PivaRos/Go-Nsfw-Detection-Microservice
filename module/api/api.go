package api

import (
	"log"
	"net/http"

	"github.com/pivaros/go-image-recognition/api/middleware"
	"github.com/pivaros/go-image-recognition/utils"
)

func Run() error {
	//init the env variables
	env, err := utils.InitEnv()
	if err != nil {
		return err
	}
	Stack := middleware.CreateStack()
	mainRouter := http.NewServeMux()
	//create server instance
	app := http.Server{
		Addr:    ":" + env.PORT,
		Handler: Stack(mainRouter),
	}
	//start listening for traffic
	log.Println("starting server on port " + env.PORT)
	err = app.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
