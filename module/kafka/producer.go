package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

func ConfigureProducer() (func(message *sarama.ProducerMessage), error) {
	// Configure the Kafka producer
	log.Println("Configuring Kafka Producer...")
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// Create a new Kafka producer
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		return nil, err
	}
	defer producer.AsyncClose()
	log.Println("Finish configuring Kafka Producer")
	return func(message *sarama.ProducerMessage) {
		producer.Input() <- message
	}, nil
}
