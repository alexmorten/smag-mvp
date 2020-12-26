package config

// EsConfig holds the configuration for elasticsearch
type EsConfig struct {
	Hosts                   []string `yaml:"hosts"`
	BulkChunkSize           int      `yaml:"bulk_chunk_size"`
	BulkFetchTimeoutSeconds int      `yaml:"bulk_fetch_timeout_seconds"`
}
