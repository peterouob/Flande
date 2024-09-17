package producer

import (
	"ecomm/config"
	"errors"
	"github.com/IBM/sarama"
	"log"
)

func NewProducer() sarama.SyncProducer {
	producer, err := sarama.NewSyncProducer(config.Config.GetStringSlice("kafka.server"), nil)
	if err != nil {
		log.Fatalf("Failed to start Kafka producer: %v", err)
	}
	log.Println("Init the kafka producer ...")
	return producer
}

func ProducerMessage(producer sarama.SyncProducer, topic, message string) error {
	producerMessage := &sarama.ProducerMessage{Topic: topic, Value: sarama.StringEncoder(message)}
	partition, offset, err := producer.SendMessage(producerMessage)
	if err != nil {
		return errors.New("failed to send message: " + err.Error())
	}
	log.Printf("Message sent to partition %d at offset %d", partition, offset)
	return nil
}
