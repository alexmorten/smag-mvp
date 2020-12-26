package config

// PostgresConfig holds all the configurable variables for Postgres
type PostgresConfig struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
}
