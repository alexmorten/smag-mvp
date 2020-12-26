package config

import (
	"io/ioutil"

	"github.com/alexmorten/smag-mvp/utils"
	"gopkg.in/yaml.v2"
)

// Config holds all of the common configuration that is parsed from config.yml
type Config struct {
	S3            S3Config         `yaml:"s3"`
	Kafka         KafkaConfig      `yaml:"kafka"`
	Postgres      PostgresConfig   `yaml:"postgres"`
	Imgproxy      ImgproxyConfig   `yaml:"imgproxy"`
	Elasticsearch EsConfig         `yaml:"elasticsearch"`
	Recognizer    RecognizerConfig `yaml:"recognizer"`
}

// LoadConfig from  config.yml as default. Can be overriden with CONFIG_PATH.
func LoadConfig() (*Config, error) {
	path := utils.GetStringFromEnvWithDefault("CONFIG_PATH", "config.yml")
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	err = yaml.Unmarshal(b, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
