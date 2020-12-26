package main

import (
	"github.com/alexmorten/smag-mvp/config"
	inserter "github.com/alexmorten/smag-mvp/insta/inserter/posts_face"
	"github.com/alexmorten/smag-mvp/kafka"
	"github.com/alexmorten/smag-mvp/service"
	"github.com/alexmorten/smag-mvp/utils"
)

func main() {
	conf, err := config.LoadConfig()
	utils.MustBeNil(err)

	qReader := kafka.NewReader(kafka.TopicNameReconResults, "faces_inserter", conf.Kafka)

	i := inserter.New(
		conf.Postgres,
		qReader,
	)

	service.CloseOnSignal(i)
	waitUntilClosed := i.Start()

	waitUntilClosed()
}
