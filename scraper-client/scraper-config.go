package client

import "github.com/alexmorten/smag-mvp/utils"

// ScraperConfig ...
type ScraperConfig struct {
	RequestTimeout    int
	RequestRetryCount int
}

// GetScraperConfig ...
func GetScraperConfig() *ScraperConfig {
	return &ScraperConfig{
		RequestTimeout:    utils.GetNumberFromEnvWithDefault("REQUEST_TIMEOUT", 1000),
		RequestRetryCount: utils.GetNumberFromEnvWithDefault("REQUEST_RETRY_COUNT", 3),
	}
}
