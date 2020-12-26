package main

import (
	"github.com/alexmorten/smag-mvp/config"
	"github.com/alexmorten/smag-mvp/db"
	inserter "github.com/alexmorten/smag-mvp/insta/inserter/posts"
	"github.com/alexmorten/smag-mvp/kafka"
	"github.com/alexmorten/smag-mvp/service"
	"github.com/alexmorten/smag-mvp/utils"
)

func main() {
	conf, err := config.LoadConfig()
	utils.MustBeNil(err)

	dbClient, err := db.NewConn(conf.Postgres)
	utils.MustBeNil(err)

	i := inserter.New(dbClient, kafka.NewReader(kafka.TopicNameScrapedPosts, "posts_inserter", conf.Kafka))

	service.CloseOnSignal(i)
	waitUntilClosed := i.Start()

	waitUntilClosed()
}
