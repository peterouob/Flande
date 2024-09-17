package service

import (
	"ecomm/kafka/producer"
	"github.com/IBM/sarama"
)

var p sarama.SyncProducer

func init() {
	p = producer.NewProducer()
}
