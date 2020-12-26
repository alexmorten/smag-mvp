package main

import (
	"github.com/alexmorten/smag-mvp/config"
	downloader "github.com/alexmorten/smag-mvp/insta/pics-downloader"
	"github.com/alexmorten/smag-mvp/kafka"
	"github.com/alexmorten/smag-mvp/service"
	"github.com/alexmorten/smag-mvp/utils"
)

func main() {
	conf, err := config.LoadConfig()
	utils.MustBeNil(err)

	qReader := kafka.NewReader(kafka.TopicNamePictureDownloads, "picture_downloader", conf.Kafka)

	i := downloader.New(qReader, &conf.S3, &conf.Postgres)

	service.CloseOnSignal(i)
	waitUntilClosed := i.Start()

	waitUntilClosed()
}
