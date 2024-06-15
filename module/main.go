package main

import (
	"log"
	"sync"

	_ "github.com/lib/pq"
	"github.com/pivaros/go-image-recognition/api"
	"github.com/pivaros/go-image-recognition/kafka"
	"github.com/pivaros/go-image-recognition/utils"
)

func main() {
	var wg sync.WaitGroup
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
	wg.Add(1)
	go kafka.Run(AppState, &wg)
	wg.Wait()
	err = api.Run(AppState)
	if err != nil {
		log.Panicln(err)
	}

}
