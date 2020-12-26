package main

import (
	"github.com/alexmorten/smag-mvp/config"
	scraper "github.com/alexmorten/smag-mvp/insta/scraper/user"
	"github.com/alexmorten/smag-mvp/kafka"
	"github.com/alexmorten/smag-mvp/service"
	"github.com/alexmorten/smag-mvp/utils"
)

func main() {
	conf, err := config.LoadConfig()
	utils.MustBeNil(err)

	s := scraper.New(
		kafka.NewReader(kafka.TopicNameUserNames, "user_scraper", conf.Kafka),
		kafka.NewWriter(kafka.TopicNameUserInfos, conf.Kafka),
		kafka.NewWriter(kafka.TopicNameUserScrapeErrors, conf.Kafka),
	)

	service.CloseOnSignal(s)
	waitUntilDone := s.Start()

	waitUntilDone()
}
