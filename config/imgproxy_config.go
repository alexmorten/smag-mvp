package config

// ImgproxyConfig holds all the configurable variables for the image proxy
type ImgproxyConfig struct {
	Address string `yaml:"address"`
	Key     string `yaml:"key"`
	Salt    string `yaml:"salt"`
}
