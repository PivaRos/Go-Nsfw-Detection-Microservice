package handlers

import "log"

func HandleEvent(topic string, message []byte) {
	switch topic {
	case "user_creation":
		HandleUserCreation(message)
	case "user_update":
		HandleUserUpdate(message)
	default:
		log.Printf("Received message from unknown topic %s: %s\n", topic, string(message))
	}
}
