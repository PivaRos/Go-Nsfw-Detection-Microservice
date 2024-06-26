package kafka

import (
	"log"
	"sync"

	"github.com/IBM/sarama"
	"github.com/pivaros/go-image-recognition/constants"
	"github.com/pivaros/go-image-recognition/kafka/handlers"
	"github.com/pivaros/go-image-recognition/utils"
)

func ConfigureConsumer(appState *utils.AppState) {
	log.Println("Configuring Kafka Consumer...")
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create a new Kafka consumer
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to start Kafka consumer: %v", err)
	}
	defer consumer.Close()

	// List of topics to consume
	topics := constants.Topics

	// WaitGroup to wait for all consumers to finish
	var wg sync.WaitGroup

	// Start a consumer for each topic
	for _, topic := range topics {
		wg.Add(1)
		go consumeTopic(consumer, topic, &wg, appState)
	}

	log.Println("Finish configuring Kafka Consumer")
	// Wait for all consumers to finish
	wg.Wait()
}

func consumeTopic(consumer sarama.Consumer, topic string, wg *sync.WaitGroup, appState *utils.AppState) {
	defer wg.Done()

	// Consume messages from the Kafka topic
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to start Kafka partition consumer for topic %s: %v", topic, err)
	}
	defer partitionConsumer.Close()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			handlers.HandleEvent(topic, msg.Value, appState)
		case err := <-partitionConsumer.Errors():
			log.Printf("Failed to consume message from %s: %v\n", topic, err)
		}
	}
}
