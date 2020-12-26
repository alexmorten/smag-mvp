package main

import (
	"github.com/alexmorten/smag-mvp/config"
	"github.com/alexmorten/smag-mvp/db"
	inserter "github.com/alexmorten/smag-mvp/insta/inserter/likes"
	"github.com/alexmorten/smag-mvp/kafka"
	"github.com/alexmorten/smag-mvp/service"
	"github.com/alexmorten/smag-mvp/utils"
)

func main() {
	conf, err := config.LoadConfig()
	utils.MustBeNil(err)

	dbClient, err := db.NewConn(conf.Postgres)
	utils.MustBeNil(err)

	s := inserter.New(dbClient, kafka.NewReader(kafka.TopicNameScrapedLikes, "like_inserter", conf.Kafka))

	service.CloseOnSignal(s)
	waitUntilClosed := s.Start()

	waitUntilClosed()
}
