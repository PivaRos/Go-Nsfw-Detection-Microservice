package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

func ConfigureProducer() func(message *sarama.ProducerMessage) {
	// Configure the Kafka producer
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// Create a new Kafka producer
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to start Kafka producer: %v", err)
	}
	defer producer.AsyncClose()
	return func(message *sarama.ProducerMessage) {
		producer.Input() <- message
	}
}
