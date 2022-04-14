package kafka

import (
	"codePix/env"
	"fmt"
	cKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
)

func NewKafkaProducer() *cKafka.Producer {
	configMap := &cKafka.ConfigMap{
		"bootstrap.servers": os.Getenv(env.KAFKA_BOOTSTRAP_SERVER),
	}
	p, err := cKafka.NewProducer(configMap)
	if err != nil {
		panic(err)
	}
	return p
}

func Publish(msg string, topic string, producer *cKafka.Producer, deliveryChan chan cKafka.Event) error {
	message := &cKafka.Message{
		TopicPartition: cKafka.TopicPartition{Topic: &topic, Partition: cKafka.PartitionAny},
		Value:          []byte(msg),
	}
	err := producer.Produce(message, deliveryChan)
	if err != nil {
		return err
	}
	return nil
}

func DeliveryReport(deliveryChan chan cKafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *cKafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Delivery fail: ", ev.TopicPartition)
			}
			fmt.Println("Delivery message to: ", ev.TopicPartition)
		}
	}
}
