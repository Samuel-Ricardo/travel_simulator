package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)
type KafkaConsumer struct {
  MessageChannel chan *kafka.Message
}

func NewKafkaConsumer(messageChannel chan *kafka.Message) *KafkaConsumer { 
  return &KafkaConsumer { MessageChannel: messageChannel }
}

