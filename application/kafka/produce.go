package kafka

import (
	"encoding/json"
	. "log"
	"os"
	"time"

	"github.com/Samuel-Ricardo/travel_simulator/application/route"
	"github.com/Samuel-Ricardo/travel_simulator/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func Produce(msg *ckafka.Message) {
  
  Println("PRODUCING")

  producer := kafka.NewKafkaProuducer()
  route := route.NewRoute()

  json.Unmarshal(msg.Value, &route)
  error := route.LoadPositions()
  if error != nil {Println(error.Error())}

  positions, err := route.ExportJsonPositions()
  if err != nil { Println(err.Error()) }

  for _, position := range positions {
    kafka.Publish(position, os.Getenv("kafkaProduceTopic"), producer)
    time.Sleep(time.Millisecond * 250)
  }
  
}
