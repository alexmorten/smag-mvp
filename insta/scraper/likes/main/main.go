package main

import (
	"github.com/alexmorten/smag-mvp/config"
	scraper "github.com/alexmorten/smag-mvp/insta/scraper/likes"
	"github.com/alexmorten/smag-mvp/kafka"
	client "github.com/alexmorten/smag-mvp/scraper-client"
	"github.com/alexmorten/smag-mvp/service"
	"github.com/alexmorten/smag-mvp/utils"
)

func main() {
	conf, err := config.LoadConfig()
	utils.MustBeNil(err)

	commentLimit := utils.GetNumberFromEnvWithDefault("LIKE_LIMIT", 24)
	config := client.GetScraperConfig()
	s := scraper.New(
		config,
		kafka.NewReader(kafka.TopicNameScrapedPosts, "like_scraper", conf.Kafka),
		kafka.NewWriter(kafka.TopicNameScrapedLikes, conf.Kafka),
		kafka.NewWriter(kafka.TopicNameLikeScrapeErrors, conf.Kafka),
		commentLimit,
	)

	service.CloseOnSignal(s)
	waitUntilClosed := s.Start()

	waitUntilClosed()
}
