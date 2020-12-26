package main

import (
	"encoding/json"

	"github.com/alexmorten/smag-mvp/config"
	"github.com/alexmorten/smag-mvp/insta/models"
	"github.com/alexmorten/smag-mvp/kafka"
	"github.com/alexmorten/smag-mvp/kafka/changestream"
	"github.com/alexmorten/smag-mvp/service"
	"github.com/alexmorten/smag-mvp/utils"

	segkafka "github.com/segmentio/kafka-go"
)

func main() {
	conf, err := config.LoadConfig()
	utils.MustBeNil(err)

	f := changestream.NewFilter(
		conf.Kafka,
		"post_face_recon_filter",
		kafka.TopicNamePGPosts,
		kafka.TopicNameReconJobs,
		filterChange)

	service.CloseOnSignal(f)
	waitUntilClosed := f.Start()

	waitUntilClosed()
}

type post struct {
	ID                 int    `json:"id"`
	InternalPictureURL string `json:"internal_picture_url"`
}

func filterChange(m *changestream.ChangeMessage) ([]segkafka.Message, error) {
	if !(m.Payload.Op == "c" || m.Payload.Op == "u") {
		return nil, nil
	}

	currentVersion := &post{}
	err := json.Unmarshal(m.Payload.After, currentVersion)
	if err != nil {
		return nil, err
	}

	if m.Payload.Op == "c" {
		return constructDownloadJobMessage(currentVersion)
	}

	previousVersion := &post{}
	err = json.Unmarshal(m.Payload.Before, previousVersion)
	if err != nil {
		return nil, err
	}

	if currentVersion.InternalPictureURL != previousVersion.InternalPictureURL {
		return constructDownloadJobMessage(currentVersion)
	}

	return nil, nil
}

func constructDownloadJobMessage(p *post) ([]segkafka.Message, error) {
	if p.InternalPictureURL == "" {
		return nil, nil
	}

	job := &models.PostDownloadJob{
		PostID:     p.ID,
		PictureURL: p.InternalPictureURL,
	}
	b, err := json.Marshal(job)
	if err != nil {
		return nil, err
	}

	return []segkafka.Message{
		{Value: b},
	}, nil
}
