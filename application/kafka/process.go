package kafka

import (
	"codePix/application/dto"
	"codePix/application/factory"
	"codePix/application/useCase"
	"codePix/domain/model"
	"codePix/env"
	"errors"
	"fmt"
	cKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jinzhu/gorm"
	"os"
)

const (
	BANK_SUFIX string = "bank"
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

func (k *KafkaProcessor) Consume() error {
	configMap := &cKafka.ConfigMap{
		"bootstrap.servers": os.Getenv(env.KAFKA_BOOTSTRAP_SERVER),
		"group.id":          os.Getenv(env.KAFKA_CONSUMER_GROUP_ID),
		"auto.offset.reset": "earliest",
	}
	consumer, err := cKafka.NewConsumer(configMap)
	if err != nil {
		panic(err)
	}

	topics := []string{os.Getenv(env.KAFKA_TRANSACTION_TOPIC), os.Getenv(env.KAFKA_TRANSACTION_CONFIRMATION_TOPIC)}
	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		return err
	}

	fmt.Println("Kafka Consumer has been started")
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			err = k.processMessage(msg)
			if err != nil {
				return err
			}
		}
	}
}

func (k *KafkaProcessor) processMessage(msg *cKafka.Message) error {
	transactionsTopics := "transactions"
	transactionsConfirmationTopics := "transactions_confirmation"

	switch topic := *msg.TopicPartition.Topic; topic {
	case transactionsTopics:
		err := k.processTransaction(msg)
		if err != nil {
			return err
		}
	case transactionsConfirmationTopics:
		err := k.processTransactionConfirmation(msg)
		if err != nil {
			return err
		}
	default:
		fmt.Println("not a valid topic", string(msg.Value))
		return errors.New("not a valid topic to process")
	}
	return nil
}

func (k *KafkaProcessor) processTransaction(msg *cKafka.Message) error {
	transaction, err := toTransaction(msg)
	if err != nil {
		return err
	}
	transactionUseCase := factory.TransactionUseCaseFactory(k.Db)

	createdTransaction, err := transactionUseCase.RegisterTransaction(
		transaction.AccountId,
		transaction.Amount,
		transaction.PixKeyIdTo,
		transaction.PixKeyKindTo,
		transaction.Description,
	)
	if err != nil {
		fmt.Println("error registering Transaction", err)
		return err
	}

	kafkaTopic := BANK_SUFIX + createdTransaction.PixKeyTo.Account.Bank.Code
	transaction.Id = createdTransaction.Id
	transaction.Status = model.TRANSACTION_PENDING
	transactionJson, err := transaction.ToJson()
	if err != nil {
		return err
	}

	err = Publish(string(transactionJson), kafkaTopic, k.Producer, k.DeliveryChan)
	if err != nil {
		return err
	}
	return nil
}

func (k *KafkaProcessor) processTransactionConfirmation(msg *cKafka.Message) error {
	transaction, err := toTransaction(msg)
	if err != nil {
		return err
	}
	transactionUseCase := factory.TransactionUseCaseFactory(k.Db)

	if transaction.Status == model.TRANSACTION_CONFIRMED {
		err = k.confirmTransaction(transaction, transactionUseCase)
		if err != nil {
			return err
		}
	} else if transaction.Status == model.TRANSACTION_COMPLETED {
		_, err := transactionUseCase.CompleteTransaction(transaction.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (k *KafkaProcessor) confirmTransaction(transaction *dto.TransactionDTO, transactionUseCase useCase.TransactionUseCase) error {
	confirmedTransaction, err := transactionUseCase.ConfirmTransaction(transaction.Id)
	if err != nil {
		return err
	}

	kafkaTopic := BANK_SUFIX + confirmedTransaction.AccountFrom.Bank.Code
	transactionJson, err := transaction.ToJson()
	if err != nil {
		return err
	}

	err = Publish(string(transactionJson), kafkaTopic, k.Producer, k.DeliveryChan)
	if err != nil {
		return err
	}
	return nil
}

func toTransaction(msg *cKafka.Message) (*dto.TransactionDTO, error) {
	transaction := dto.NewTransaction()
	err := transaction.ParseJson(msg.Value)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}
