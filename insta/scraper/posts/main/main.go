package main

import (
	scraper "github.com/alexmorten/smag-mvp/insta/scraper/posts"
	"github.com/alexmorten/smag-mvp/kafka"
	client "github.com/alexmorten/smag-mvp/scraper-client"
	"github.com/alexmorten/smag-mvp/service"
	"github.com/alexmorten/smag-mvp/utils"
)

func main() {
	awsServiceAddress := utils.GetStringFromEnvWithDefault("AWS_SERVICE_ADDRESS", "")
	nameReaderConfig, infoWriterConfig, errWriterConfig := kafka.GetInstaPostsScraperConfig()

	config := client.GetScraperConfig()
	s := scraper.New(config, awsServiceAddress, kafka.NewReader(nameReaderConfig), kafka.NewWriter(infoWriterConfig), kafka.NewWriter(errWriterConfig))

	service.CloseOnSignal(s)
	waitUntilClosed := s.Start()

	waitUntilClosed()
}
