package cmd

import (
	"codePix/application/grpc"
	"codePix/application/kafka"
	db "codePix/config"
	"codePix/env"
	cKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/cobra"
	"os"
)

var gRpcPortNumber int

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Run gRPC and Kafka Consumer together",
	Run: func(cmd *cobra.Command, args []string) {
		// gRPC
		database := db.ConnectDB(os.Getenv(env.CURRENT_ENV))
		go grpc.StartGrpcServer(database, gRpcPortNumber)

		// Kafka
		producer := kafka.NewKafkaProducer()
		deliveryChan := make(chan cKafka.Event)
		go kafka.DeliveryReport(deliveryChan)
		kafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChan)
		kafkaProcessor.Consume()
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
	allCmd.Flags().IntVarP(&gRpcPortNumber, "grpc-port", "p", 50051, "gRPC Port")
}
