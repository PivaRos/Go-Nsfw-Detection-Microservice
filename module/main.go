package main

import (
	"log"

	"github.com/pivaros/go-image-recognition/api"
	"github.com/pivaros/go-image-recognition/kafka"
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

	go kafka.Run()
	err := api.Run()
	if err != nil {
		log.Panicln(err)
	}

}
