package main

import (
	"context"
	"os"
	"time"

	"github.com/alexmorten/smag-mvp/utils"
	"github.com/segmentio/kafka-go"
)

func main() {
	kafkaAddress := utils.GetStringFromEnvWithDefault("KAFKA_ADDRESS", "my-kafka:9092")
	topic := utils.GetStringFromEnvWithDefault("KAFKA_INSTAGRAM_TOPIC", "user_names")

	if len(os.Args) < 2 {
		panic("Invalid argumemts. Usage: cli <username>")
	}

	userNameArg := os.Args[1]

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaAddress},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer w.Close()
	t, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err := w.WriteMessages(t, kafka.Message{
		Value: []byte(userNameArg),
	})
	utils.PanicIfNotNil(err)
}
