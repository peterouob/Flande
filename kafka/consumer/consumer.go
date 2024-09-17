package consumer

import (
	"ecomm/config"
	"errors"
	"github.com/IBM/sarama"
	"log"
	"os"
	"os/signal"
	"sync"
)

var partitionConsumer sarama.PartitionConsumer

func StartConsumer(topic string, wg *sync.WaitGroup) error {
	defer wg.Done()
	consumer, err := sarama.NewConsumer(config.Config.GetStringSlice("kafka.server"), nil)
	if err != nil {
		return errors.New("error to init new consumer: " + err.Error())
	}
	defer consumer.Close()
	log.Println("Init the kafka consumer ...")

	partitions, err := consumer.Partitions(topic)
	if err != nil {
		return errors.New("error occur consumer.Partitions: " + err.Error())
	}
	for _, partition := range partitions {
		partitionConsumer, err = consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			return errors.New("error occur consumer.ConsumePartition: " + err.Error())
		}
		go func() {
			for {
				select {
				case msg := <-partitionConsumer.Messages():
					log.Printf("Consumed message: %s\n", string(msg.Value))
				case err := <-partitionConsumer.Errors():
					log.Printf("error occur in partitionConsumer: %s", err.Error())
				}
			}
		}()
	}
	defer partitionConsumer.Close()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	<-sigs
	log.Println("shut down consumer ...")
	return nil
}
