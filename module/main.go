package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/pivaros/go-image-recognition/api"
	"github.com/pivaros/go-image-recognition/kafka"
	"github.com/pivaros/go-image-recognition/utils"
)

func main() {
	// the highest panic catch
	defer func() {
		if r := recover(); r != nil {
			log.Println("api: top level panic:", r)
		} else {
			log.Println("api: couldn't recover panic in top level")
		}
	}()
	//setup env
	env, err := utils.InitEnv()
	if err != nil {
		log.Panicln(err)
	}

	//setup postgres db
	db, err := utils.ConnectPostgres(env.Connection_String)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Successfully connected to PostgreSQL")
	var AppState = &utils.AppState{
		Env: env,
		Db:  db,
	}
	//run the applications and pass the postgres connection
	go kafka.Run(AppState)
	err = api.Run(AppState)
	if err != nil {
		log.Panicln(err)
	}

}
