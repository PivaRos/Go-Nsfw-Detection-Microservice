package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

func ConfigureConsumer() {
	log.Println("Configuring Kafka Consumer...")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create a new Kafka consumer
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to start Kafka consumer: %v", err)
	}
	defer consumer.Close()

	// Consume messages from the Kafka topic
	topic := "test1"
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to start Kafka partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message: %s\n", string(msg.Value))
		case err := <-partitionConsumer.Errors():
			log.Printf("Failed to consume message: %v\n", err)
		}
	}

}
