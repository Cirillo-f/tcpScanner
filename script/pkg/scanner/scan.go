package scanner

import (
	"context"
	"sort"
)

type Scanner struct {
	config *Config
}

func New(opts ...Option) *Scanner {
	return &Scanner{
		config: NewConfig(opts...),
	}
}

func (s *Scanner) Scan(ctx context.Context, host string, ports []int, progress chan<- int) ([]int, error) {
	resolvedHost, err := ResolveHost(host)
	if err != nil {
		return nil, err
	}

	pool := newWorkerPool(ctx, resolvedHost, ports, s.config, progress)
	openPorts := pool.run()

	sort.Ints(openPorts)
	return openPorts, nil
}

func (s *Scanner) ScanPopular(ctx context.Context, host string, progress chan<- int) ([]int, error) {
	cfg := NewConfigForPopular()
	originalConfig := s.config
	s.config = cfg
	defer func() { s.config = originalConfig }()
	return s.Scan(ctx, host, PopularPorts, progress)
}

func (s *Scanner) ScanAll(ctx context.Context, host string, progress chan<- int) ([]int, error) {
	return s.Scan(ctx, host, AllPorts(), progress)
}
