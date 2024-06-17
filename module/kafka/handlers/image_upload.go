package handlers

import (
	"encoding/json"
	"log"

	"github.com/pivaros/go-image-recognition/structs"
	"github.com/pivaros/go-image-recognition/utils"
)

func HandleImageUpload(messageRaw []byte, appState *utils.AppState) {
	//panic handler
	defer func() {
		if r := recover(); r != nil {
			log.Println("kafka thread: top level panic:", r)
		}
	}()
	//put this message (image) through the image classification model
	//and then update the db
	var message structs.ImageUploadMessage
	var err = json.Unmarshal(messageRaw, &message)
	if err != nil {
		panic(err)
	}
	results, err := appState.ClassificationFunc(message.Base64)
	if err != nil {
		panic(err)
	}

	log.Println(*results)

}
