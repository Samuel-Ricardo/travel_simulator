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

func Publish (msg string, topic string, producer *kafka.Producer) error {
  message := &kafka.Message{
    TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
    Value:          []byte(msg),
  }

  err := producer.Produce(message, nil)
  if err != nil { return err }
  
  return nil
}
