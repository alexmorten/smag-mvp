package config

// S3Config holds all the configurable variables for S3
type S3Config struct {
	PictureBucketName string `yaml:"picture_bucket_name"`
	Region            string `yaml:"region"`
	Endpoint          string `yaml:"endpoint"`
	AccessKeyID       string `yaml:"key_id"`
	SecretAccessKey   string `yaml:"secret"`
	UseSSL            bool   `yaml:"use_ssl"`
}
