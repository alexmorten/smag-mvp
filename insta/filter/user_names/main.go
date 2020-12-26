package main

import (
	"encoding/json"

	"github.com/alexmorten/smag-mvp/config"
	ikafka "github.com/alexmorten/smag-mvp/kafka"
	"github.com/alexmorten/smag-mvp/kafka/changestream"
	"github.com/alexmorten/smag-mvp/service"
	"github.com/alexmorten/smag-mvp/utils"
	"github.com/segmentio/kafka-go"
)

func main() {

	conf, err := config.LoadConfig()
	utils.MustBeNil(err)

	f := changestream.NewFilter(
		conf.Kafka,
		"user_names_filter",
		ikafka.TopicNamePGUsers,
		ikafka.TopicNameUserNames,
		filterChange)

	service.CloseOnSignal(f)
	waitUntilClosed := f.Start()

	waitUntilClosed()
}

type user struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
}

func filterChange(m *changestream.ChangeMessage) ([]kafka.Message, error) {
	if m.Payload.Op != "c" {
		return nil, nil
	}

	u := &user{}
	err := json.Unmarshal(m.Payload.After, u)
	if err != nil {
		return nil, err
	}

	return []kafka.Message{{Value: []byte(u.UserName)}}, nil
}
