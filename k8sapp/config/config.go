package config

import (
	"goinaction/k8sapp/logger"

	"github.com/kelseyhightower/envconfig"
)

const (
	SERVICENAME = "K8SAPP"
)

type Config struct {
	LocalHost string       `split_words:"true"`
	LocalPort int          `split_words:"true"`
	LogLevel  logger.Level `split_words:"true"`
}

func (c *Config) Load(serviceName string) error {
	return envconfig.Process(serviceName, c)
}
