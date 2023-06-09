package main

import (
	"fmt"
	"log"

	kafka2 "github.com/Samuel-Ricardo/travel_simulator/application/kafka"
	"github.com/Samuel-Ricardo/travel_simulator/infra/kafka"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)


func init() {
  err := godotenv.Load()
  if err != nil { log.Fatal("Error on Loading .env File") }
}

func main() {

  messageChannel := make(chan *ckafka.Message)
  consumer := kafka.NewKafkaConsumer(messageChannel)

  go consumer.Consume()

  for message := range messageChannel{
    fmt.Println(string(message.Value)) 
    go kafka2.Produce(message)
  }
}
