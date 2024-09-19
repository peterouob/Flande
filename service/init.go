package service

import (
	"ecomm/kafka/producer"
	"github.com/IBM/sarama"
)

var Producer sarama.SyncProducer

func init() {
	Producer = producer.NewProducer()
}
