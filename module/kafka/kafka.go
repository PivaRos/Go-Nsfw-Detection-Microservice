package kafka

import (
	"log"
	"sync"

	"github.com/pivaros/go-image-recognition/kafka/nsfw"
	"github.com/pivaros/go-image-recognition/utils"
)

func Run(appState *utils.AppState, wg *sync.WaitGroup) {
	// the highest panic catch
	defer func() {
		if r := recover(); r != nil {
			log.Println("Kafka thread: top level panic:", r)
		}
	}()
	//configure kafka producer
	producer, err := ConfigureProducer()
	if err != nil {
		panic(err)
	}
	appState.ProduceMessage = producer
	//configure the python nsfw model
	classificationFunc, err := nsfw.ConfigureModel()
	if err != nil {
		panic(err)
	}
	appState.ClassificationFunc = classificationFunc

	//resume api
	wg.Done()
	//configure kafka consumer
	ConfigureConsumer(appState)

}
