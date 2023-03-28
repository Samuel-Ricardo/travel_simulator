package kafka

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafkaProuducer() *kafka.Producer {
  configMap := &kafka.ConfigMap{
    "bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
  }

  producer, err := kafka.NewProducer(configMap)
  if err != nil { log.Println(err.Error()) }

  return producer;
}
