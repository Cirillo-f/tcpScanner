package scanner

import (
	"net"
	"time"
)

const (
	maxPorts       = 65535
	defaultTimeout = 1 * time.Second
	defaultWorkers = 500
	popularTimeout = 500 * time.Millisecond
)

type Config struct {
	Timeout time.Duration
	Workers int
	Dialer  *net.Dialer
}

type Option func(*Config)

func WithTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.Timeout = timeout
	}
}

func WithWorkers(workers int) Option {
	return func(c *Config) {
		c.Workers = workers
	}
}

func NewConfig(opts ...Option) *Config {
	cfg := &Config{
		Timeout: defaultTimeout,
		Workers: defaultWorkers,
		Dialer: &net.Dialer{
			Timeout: defaultTimeout,
		},
	}

	for _, opt := range opts {
		opt(cfg)
	}

	cfg.Dialer.Timeout = cfg.Timeout
	return cfg
}

func NewConfigForPopular(opts ...Option) *Config {
	cfg := NewConfig(opts...)
	cfg.Timeout = popularTimeout
	cfg.Dialer.Timeout = popularTimeout
	return cfg
}
