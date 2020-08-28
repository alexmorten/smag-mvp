package main

import (
	inserter "github.com/alexmorten/smag-mvp/insta/inserter/likes"
	"github.com/alexmorten/smag-mvp/kafka"
	"github.com/alexmorten/smag-mvp/service"
	"github.com/alexmorten/smag-mvp/utils"
)

func main() {
	postgresHost := utils.GetStringFromEnvWithDefault("POSTGRES_HOST", "127.0.0.1")
	postgresPassword := utils.GetStringFromEnvWithDefault("POSTGRES_PASSWORD", "")

	qReaderConfig := kafka.GetInserterConfig()

	s := inserter.New(postgresHost, postgresPassword, kafka.NewReader(qReaderConfig))

	service.CloseOnSignal(s)
	waitUntilClosed := s.Start()

	waitUntilClosed()
}
