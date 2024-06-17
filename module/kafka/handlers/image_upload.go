package handlers

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
	"github.com/pivaros/go-image-recognition/kafka/nsfw"
	"github.com/pivaros/go-image-recognition/structs"
	"github.com/pivaros/go-image-recognition/utils"
)

type CompositeStruct struct {
	ImageGuid            string                  `json:"imageGuid"`
	ClassificationResult *[]nsfw.DetectionResult `json:"classificationResult"`
}

func HandleImageUpload(messageRaw []byte, appState *utils.AppState) {
	//panic handler
	defer func() {
		if r := recover(); r != nil {
			log.Println("kafka thread: panic recovery HandleImageUpload", r)
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

	classificationResultsJSON, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}
	kafkaMessageRaw, err := json.Marshal(CompositeStruct{
		ImageGuid:            message.ImageGuid,
		ClassificationResult: results,
	})
	if err != nil {
		panic(err)
	}
	query := `INSERT INTO image_classifications (image_guid, classification_result) VALUES ($1, $2)`
	_, err = appState.Db.Exec(query, message.ImageGuid, classificationResultsJSON)
	if err != nil {
		panic(err)
	}
	producedMessage := &sarama.ProducerMessage{
		Topic: "image_classified",
		Value: sarama.ByteEncoder(kafkaMessageRaw),
	}

	appState.ProduceMessage(producedMessage)

}
