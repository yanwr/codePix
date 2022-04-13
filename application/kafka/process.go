package kafka

import (
	"fmt"
	cKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jinzhu/gorm"
)

type KafkaProcessor struct {
	Db           *gorm.DB
	Producer     *cKafka.Producer
	DeliveryChan chan cKafka.Event
}

func NewKafkaProcessor(db *gorm.DB, producer *cKafka.Producer, deliveryChan chan cKafka.Event) *KafkaProcessor {
	return &KafkaProcessor{
		Db:           db,
		Producer:     producer,
		DeliveryChan: deliveryChan,
	}
}

func (k *KafkaProcessor) Consume() {
	configMap := &cKafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "consumergroup",
		"auto.offset.reset": "earliest",
	}
	consumer, err := cKafka.NewConsumer(configMap)
	if err != nil {
		panic(err)
	}

	topics := []string{"test"}
	consumer.SubscribeTopics(topics, nil)

	fmt.Println("Kafka Consumer has been started")
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(msg.Value))
		}
	}
}
