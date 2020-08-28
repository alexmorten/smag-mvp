package main

import (
	inserter "github.com/alexmorten/smag-mvp/insta/inserter/posts"
	"github.com/alexmorten/smag-mvp/kafka"
	"github.com/alexmorten/smag-mvp/service"
	"github.com/alexmorten/smag-mvp/utils"
)

func main() {
	postgresHost := utils.GetStringFromEnvWithDefault("POSTGRES_HOST", "127.0.0.1")
	postgresPassword := utils.GetStringFromEnvWithDefault("POSTGRES_PASSWORD", "")

	kafkaAddress := utils.GetStringFromEnvWithDefault("KAFKA_ADDRESS", "127.0.0.1:9092")

	groupID := "posts_inserter"
	rTopic := kafka.TopicNameScrapedPosts
	qReaderConfig := kafka.NewReaderConfig(kafkaAddress, groupID, rTopic)

	i := inserter.New(postgresHost, postgresPassword, kafka.NewReader(qReaderConfig))

	service.CloseOnSignal(i)
	waitUntilClosed := i.Start()

	waitUntilClosed()
}
