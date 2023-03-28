package kafka

import (
  ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
  MessageChannel chan *ckafka
}
