package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/Samuel-Ricardo/travel_simulator/application/route"
	"github.com/Samuel-Ricardo/travel_simulator/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func Produce(msg *ckafka.Message) {
  
  producer := kafka.NewKafkaProuducer()
  route := route.NewRoute()

  json.Unmarshal(msg.Value, &route)
  route.LoadPositions()
  
  positions, err := route.ExportJsonPositions()
  if err != nil { log.Println(err.Error()) }
  
  for _, p := range positions {
    kafka.Publish(p, os.Getenv("kafkaProducerTopic"), producer)
    time.Sleep(time.Millisecond * 500)
  }
  
}
