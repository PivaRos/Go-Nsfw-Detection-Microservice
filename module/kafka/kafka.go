package kafka

import (
	"log"
	"sync"

	"github.com/pivaros/go-image-recognition/utils"
)

func Run(appState *utils.AppState, wg *sync.WaitGroup) {
	// the highest panic catch
	defer func() {
		if r := recover(); r != nil {
			log.Println("kafka thread: top level panic:", r)
		}
	}()
	//configure kafka producer
	appState.ProduceMessage = ConfigureProducer()
	log.Println("Finish configuring Kafka producer")
	//resume api
	wg.Done()
	//configure kafka consumer

}
