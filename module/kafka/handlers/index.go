package handlers

import (
	"log"

	"github.com/pivaros/go-image-recognition/constants"
	"github.com/pivaros/go-image-recognition/utils"
)

func HandleEvent(topic string, message []byte, appState *utils.AppState) {
	switch topic {
	case string(constants.ImageUpload):
		HandleImageUpload(message, appState)
	default:
		log.Printf("Received message from unhandled topic %s: %s\n", topic, string(message))
	}
}
