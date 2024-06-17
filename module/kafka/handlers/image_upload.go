package handlers

import (
	"log"
)

func HandleImageUpload(message []byte) {
	log.Println("image_upload", string(message))
}
