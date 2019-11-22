package insta_post_text_filter

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/codeuniversity/smag-mvp/kafka/changestream"
	"github.com/codeuniversity/smag-mvp/worker"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/segmentio/kafka-go"
	"log"
)

const (
	elasticProtocol = "http://%s"
)

type Inserter struct {
	*worker.Worker

	client  *elasticsearch.Client
	kReader *kafka.Reader

	insertFunc InserterFunc
}

type InserterFunc func(*changestream.ChangeMessage, *elasticsearch.Client) error

func New(elasticSearchAddress string, kReader *kafka.Reader, inserterFunc InserterFunc) *Inserter {
	inserter := &Inserter{}

	inserter.kReader = kReader
	inserter.insertFunc = inserterFunc

	inserter.initializeMongo(elasticSearchAddress)
	return inserter
}

func (i *Inserter) runStep() error {
	m, err := i.kReader.FetchMessage(context.Background())

	if err != nil {
		return err
	}

	changeMessage := &changestream.ChangeMessage{}

	err = json.Unmarshal(m.Value, changeMessage)

	if err != nil {
		return err
	}

	err = i.insertFunc(changeMessage, i.client)

	if err != nil {
		return err
	}

	log.Println("Inserted")
	return i.kReader.CommitMessages(context.Background(), m)
}

func (i *Inserter) initializeMongo(elasticSearchAddress string) *elasticsearch.Client {
	url := fmt.Sprintf(elasticProtocol, elasticSearchAddress)

	cfg := elasticsearch.Config{
		Addresses: []string{
			url,
		},
	}
	client, err := elasticsearch.NewClient(cfg)

	if err != nil {
		panic(err)
	}
	return client
}
