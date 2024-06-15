package kafka

import (
	"log"

	"github.com/pivaros/go-image-recognition/utils"
)

func Run(appState *utils.AppState) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("kafka thread: top level panic:", r)
		}
	}()

}
