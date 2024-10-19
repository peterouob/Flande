package service

import (
	"ecomm/etcd"
	"ecomm/kafka/producer"
	"errors"
	"github.com/IBM/sarama"
	eclient "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	gresolver "google.golang.org/grpc/resolver"
)

var Producer sarama.SyncProducer

func init() {
	Producer = producer.NewProducer()
}

func InitBuilder() (gresolver.Builder, error) {
	cli, err := eclient.NewFromURL(etcd.EtcdAddress)
	if err != nil {
		return nil, errors.New("error in client etcd from url")
	}
	builder, err := resolver.NewBuilder(cli)
	if err != nil {
		return nil, errors.New("error in builder")
	}
	return builder, nil
}
