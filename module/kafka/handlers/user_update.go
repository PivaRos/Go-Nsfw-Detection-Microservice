package handlers

import (
	"log"
)

func HandleUserUpdate(message []byte) {
	log.Println("HandleUserUpdate", string(message))
}
