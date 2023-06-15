package kafka

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafkaProuducer() *kafka.Producer {
  configMap := &kafka.ConfigMap{
    "bootstrap.servers": os.Getenv("kafkaBootstrapServer"),
    //"security.protocol": os.Getenv("security.protocol"),
		//"sasl.mechanisms":   os.Getenv("sasl.mechanisms"),
		//"sasl.username":     os.Getenv("sasl.username"),
		//"sasl.password":     os.Getenv("sasl.password"),
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
