package main

import (
	"github.com/alexmorten/smag-mvp/config"
	inserter "github.com/alexmorten/smag-mvp/insta/inserter/postgres"
	"github.com/alexmorten/smag-mvp/kafka"
	"github.com/alexmorten/smag-mvp/service"
	"github.com/alexmorten/smag-mvp/utils"
)

func main() {
	conf, err := config.LoadConfig()
	utils.MustBeNil(err)

	i := inserter.New(
		conf.Postgres,
		kafka.NewReader(kafka.TopicNameUserInfos, "user_info_inserter", conf.Kafka),
	)

	service.CloseOnSignal(i)
	waitUntilClosed := i.Start()

	waitUntilClosed()
}
