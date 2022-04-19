package env

const (
	CURRENT_ENV  = "env"
	DEBUG        = "debug"
	AUTO_MIGRATE = "AutoMigrateDb"

	DB_TYPE      = "dbType"
	DNS          = "dsn"
	DB_TYPE_TEST = "dbTypeTest"
	DNS_TEST     = "dsnTest"

	KAFKA_BOOTSTRAP_SERVER               = "kafkaBootstrapServers"
	KAFKA_CONSUMER_GROUP_ID              = "kafkaConsumerGroupId"
	KAFKA_TRANSACTION_TOPIC              = "kafkaTransactionTopic"
	KAFKA_TRANSACTION_CONFIRMATION_TOPIC = "kafkaTransactionConfirmationTopic"
)
