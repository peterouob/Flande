package service

import (
	"ecomm/config"
	"ecomm/kafka/producer"
	"ecomm/token"
	"github.com/IBM/sarama"
)

var Producer sarama.SyncProducer
var Token token.Token

func init() {
	Producer = producer.NewProducer()
	tokenKey := config.Config.GetString("token.key")
	Token = *token.NewToken(tokenKey)
}
