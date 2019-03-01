package prisma

import (
	"github.com/caarlos0/env"
)

type Config struct {
	Endpoint string `env:"PRISMA_ENDPOINT" envDefault:"http://localhost:4466/graphql/dev"`
	Secret   string `env:"PRISMA_SECRET"`
}

func NewConfig() (*Options, error) {
	config := &Config{}
	err := env.Parse(config)
	return (*Options)(config), err
}

func NewClient(config *Options) *Client {
	return New(config)
}
