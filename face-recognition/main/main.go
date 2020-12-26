package main

import (
	"github.com/alexmorten/smag-mvp/config"
	recognition "github.com/alexmorten/smag-mvp/face-recognition"
	"github.com/alexmorten/smag-mvp/kafka"
	"github.com/alexmorten/smag-mvp/service"
	"github.com/alexmorten/smag-mvp/utils"
)

func main() {
	conf, err := config.LoadConfig()
	utils.MustBeNil(err)
	qReader := kafka.NewReader(kafka.TopicNameReconJobs, "face_recognizer", conf.Kafka)
	qWriter := kafka.NewWriter(kafka.TopicNameReconResults, conf.Kafka)

	r := recognition.New(qReader, qWriter, conf)

	service.CloseOnSignal(r)
	waitUntilDone := r.Start()

	waitUntilDone()
}
