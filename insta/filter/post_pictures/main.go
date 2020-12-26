package main

import (
	"encoding/json"
	"fmt"

	"github.com/alexmorten/smag-mvp/config"
	"github.com/alexmorten/smag-mvp/insta/models"
	kf "github.com/alexmorten/smag-mvp/kafka"
	"github.com/alexmorten/smag-mvp/kafka/changestream"
	"github.com/alexmorten/smag-mvp/service"
	"github.com/alexmorten/smag-mvp/utils"

	"github.com/segmentio/kafka-go"
)

func main() {
	conf, err := config.LoadConfig()
	utils.MustBeNil(err)

	groupID := "post_pictures_filter"
	changesTopic := kf.TopicNamePGPosts
	downloadTopic := kf.TopicNamePictureDownloads

	f := changestream.NewFilter(conf.Kafka, groupID, changesTopic, downloadTopic, filterChange)
	service.CloseOnSignal(f)
	waitUntilClosed := f.Start()

	waitUntilClosed()
}

type post struct {
	ID         int    `json:"id"`
	PictureURL string `json:"picture_url"`
}

func filterChange(m *changestream.ChangeMessage) ([]kafka.Message, error) {
	fmt.Println("filtering", "payload op", m.Payload.Op)
	if !(m.Payload.Op == "r" || m.Payload.Op == "c" || m.Payload.Op == "u") {
		return nil, nil
	}

	currentVersion := &post{}
	err := json.Unmarshal(m.Payload.After, currentVersion)
	if err != nil {
		return nil, err
	}

	if currentVersion.PictureURL == "" {
		return nil, nil
	}

	if m.Payload.Op == "c" || m.Payload.Op == "r" {
		return constructDownloadJobMessage(currentVersion)
	}

	previousVersion := &post{}
	err = json.Unmarshal(m.Payload.Before, previousVersion)
	if err != nil {
		return nil, err
	}

	if currentVersion.PictureURL != previousVersion.PictureURL {
		return constructDownloadJobMessage(currentVersion)
	}

	return nil, nil
}

func constructDownloadJobMessage(p *post) ([]kafka.Message, error) {
	job := &models.PostDownloadJob{
		PostID:     p.ID,
		PictureURL: p.PictureURL,
	}
	b, err := json.Marshal(job)
	if err != nil {
		return nil, err
	}

	return []kafka.Message{
		{Value: b},
	}, nil
}
