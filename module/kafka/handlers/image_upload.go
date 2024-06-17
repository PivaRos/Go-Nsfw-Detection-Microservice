package handlers

import (
	"github.com/pivaros/go-image-recognition/utils"
)

func HandleImageUpload(message []byte, appState *utils.AppState) {
	//put this message (image) through the image classification model
	//and then update the db
}
