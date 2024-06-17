package handlers

import "log"

func HandleEvent(topic string, message []byte) {
	switch topic {
	case "image_upload":
		HandleImageUpload(message)
	default:
		log.Printf("Received message from unhandled topic %s: %s\n", topic, string(message))
	}
}
