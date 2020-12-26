package main

import (
	"github.com/alexmorten/smag-mvp/config"
	scraper "github.com/alexmorten/smag-mvp/insta/scraper/posts"
	"github.com/alexmorten/smag-mvp/kafka"
	client "github.com/alexmorten/smag-mvp/scraper-client"
	"github.com/alexmorten/smag-mvp/service"
	"github.com/alexmorten/smag-mvp/utils"
)

func main() {
	conf, err := config.LoadConfig()
	utils.MustBeNil(err)

	config := client.GetScraperConfig()
	s := scraper.New(
		config,
		kafka.NewReader(kafka.TopicNameUserNames, "post_scraper", conf.Kafka),
		kafka.NewWriter(kafka.TopicNameScrapedPosts, conf.Kafka),
		kafka.NewWriter(kafka.TopicNamePostScrapeErrors, conf.Kafka),
	)

	service.CloseOnSignal(s)
	waitUntilClosed := s.Start()

	waitUntilClosed()
}
