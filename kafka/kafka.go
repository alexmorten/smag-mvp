package kafka

import (
	"fmt"
	"time"

	"github.com/alexmorten/smag-mvp/config"
	"github.com/alexmorten/smag-mvp/utils"
	"github.com/segmentio/kafka-go"
)

// ReaderConfig is an internal helper structure for creating
// a new pre-configurated kafka-go Reader
type ReaderConfig struct {
	Address string
	Topic   string
}

// WriterConfig is an internal helper structure for creating
// a new pre-configurated kafka-go Writer
type WriterConfig struct {
	Address string
	Topic   string
	Async   bool
}

// NewWriterConfig creates a new WriterConfig structure
func NewWriterConfig(kafkaAddress, topic string, async bool) *WriterConfig {
	return &WriterConfig{
		Address: kafkaAddress,
		Topic:   topic,
		Async:   async,
	}
}

// NewReader creates a kafka-go Reader structure by using common
// configuration and additionally applying the ReaderConfig on it
func NewReader(topic string, groupID string, conf config.KafkaConfig) *kafka.Reader {
	version := utils.GetStringFromEnvWithDefault("KAFKA_GROUPID_VERSION", "1")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:               []string{conf.Address},
		GroupID:               fmt.Sprintf("%s.v%s", groupID, version),
		Topic:                 topic,
		MinBytes:              1e3,  // 1KB
		MaxBytes:              10e6, // 10MB
		QueueCapacity:         10000,
		CommitInterval:        time.Second,
		ReadBackoffMax:        time.Second * 5,
		WatchPartitionChanges: true,
		RetentionTime:         time.Hour * 24 * 30,
	})
}

// NewWriter creates a kafka-go Writer
func NewWriter(topic string, conf config.KafkaConfig) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{conf.Address},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
		Async:    true,
	})
}
