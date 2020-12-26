package main

import (
	"github.com/alexmorten/smag-mvp/config"
	downloader "github.com/alexmorten/smag-mvp/insta/pics-downloader"
	"github.com/alexmorten/smag-mvp/kafka"
	"github.com/alexmorten/smag-mvp/service"
	"github.com/alexmorten/smag-mvp/utils"
)

func main() {
	kafkaAddress := utils.GetStringFromEnvWithDefault("KAFKA_ADDRESS", "127.0.0.1:9092")
	groupID := "picture_downloader"
	jobsTopic := kafka.TopicNamePictureDownloads
	qReader := kafka.NewReader(kafka.NewReaderConfig(kafkaAddress, groupID, jobsTopic))

	s3Config := config.GetS3Config()
	postgresConfig := config.GetPostgresConfig()

	i := downloader.New(qReader, s3Config, postgresConfig)

	service.CloseOnSignal(i)
	waitUntilClosed := i.Start()

	waitUntilClosed()
}
