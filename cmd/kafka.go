package cmd

import (
	"codePix/application/kafka"
	db "codePix/config"
	"fmt"
	cKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/cobra"
	"os"
)

// kafkaCmd represents the kafka command
var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Start consuming transactions using Apache Kafka",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Producing message")
		database := db.ConnectDB(os.Getenv("env"))
		producer := kafka.NewKafkaProducer()
		deliveryChan := make(chan cKafka.Event)

		kafka.Publish("Heelllooowww, Kafka", "test", producer, deliveryChan)
		// It's going to run in Async way
		go kafka.DeliveryReport(deliveryChan)

		kafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChan)
		kafkaProcessor.Consume()
	},
}

func init() {
	rootCmd.AddCommand(kafkaCmd)
}
