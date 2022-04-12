package cmd

import (
	"codePix/application/kafka"
	"fmt"
	cKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/cobra"
)

// kafkaCmd represents the kafka command
var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Start consuming transactions using Apache Kafka",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Producing message")
		producer := kafka.NewKafkaProducer()
		deliveryChan := make(chan cKafka.Event)
		kafka.Publish("Heelllooowww, Kafka", "test", producer, deliveryChan)
		kafka.DeliveryReport(deliveryChan)
	},
}

func init() {
	rootCmd.AddCommand(kafkaCmd)
}
