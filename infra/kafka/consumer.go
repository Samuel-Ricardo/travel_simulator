package kafka

import (
	"fmt"
	"log"
	"os"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)
type KafkaConsumer struct {
  MessageChannel chan *kafka.Message
}

func NewKafkaConsumer(messageChannel chan *kafka.Message) *KafkaConsumer { 
  return &KafkaConsumer { MessageChannel: messageChannel }
}

func (k *KafkaConsumer) Consume() {
  configMap := &kafka.ConfigMap{
    "bootstrap.server": os.Getenv("kafkaBootstrapServers"),
    "group.id":         os.Getenv("kafkaConsumerGroupId"),
  }

  consumer, err := kafka.NewConsumer(configMap)
  if err != nil { log.Fatalf("Error on cosuming kafka message: "+err.Error()) }

  topics := []string{os.Getenv("kafkaReadTopic")}
  consumer.SubscribeTopics(topics, nil)
  fmt.Println("Kafka Consumer  has been started")

  for {
    msg, err := consumer.ReadMessage(-1)
    if err == nil { k.MessageChannel <- msg }
  }
}
