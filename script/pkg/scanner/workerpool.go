package scanner

import (
	"context"
	"net"
	"strconv"
	"sync"
)

type workerPool struct {
	config     *Config
	host       string
	ports      []int
	results    chan int
	progress   chan<- int
	ctx        context.Context
	cancel     context.CancelFunc
}

func newWorkerPool(ctx context.Context, host string, ports []int, cfg *Config, progress chan<- int) *workerPool {
	poolCtx, cancel := context.WithCancel(ctx)
	return &workerPool{
		config:   cfg,
		host:     host,
		ports:    ports,
		results:  make(chan int, cfg.Workers),
		progress: progress,
		ctx:      poolCtx,
		cancel:   cancel,
	}
}


func (wp *workerPool) run() []int {
	workers := wp.config.Workers
	if workers > len(wp.ports) {
		workers = len(wp.ports)
	}
	if workers <= 0 {
		workers = 1
	}

	portsCh := make(chan int, workers)
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			wp.worker(portsCh)
		}()
	}

	go func() {
		defer close(portsCh)
		for _, port := range wp.ports {
			select {
			case <-wp.ctx.Done():
				return
			case portsCh <- port:
			}
		}
	}()

	go func() {
		wg.Wait()
		close(wp.results)
	}()

	openPorts := make([]int, 0, 32)
	for {
		select {
		case <-wp.ctx.Done():
			return openPorts
		case port, ok := <-wp.results:
			if !ok {
				return openPorts
			}
			openPorts = append(openPorts, port)
		}
	}
}

func (wp *workerPool) worker(portsCh <-chan int) {
	for {
		select {
		case <-wp.ctx.Done():
			return
		case port, ok := <-portsCh:
			if !ok {
				return
			}

			addr := net.JoinHostPort(wp.host, strconv.Itoa(port))
			conn, err := wp.config.Dialer.DialContext(wp.ctx, "tcp", addr)
			if err == nil {
				conn.Close()
				select {
				case wp.results <- port:
				case <-wp.ctx.Done():
					return
				}
			}

			if wp.progress != nil {
				select {
				case wp.progress <- 1:
				default:
				}
			}
		}
	}
}
