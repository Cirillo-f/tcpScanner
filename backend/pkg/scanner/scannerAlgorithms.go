package scanner

import (
	"net"
	"strconv"
	"time"
)

const (
	MAX_PORTS = 65535
	TIMEOUT   = 1 * time.Second
	WORKERS   = 500
)

func TCPScanner(host string) []int {
	ports := make(chan int, WORKERS)
	results := make(chan int)
	var openPorts []int

	for i := 0; i < WORKERS; i++ {
		go worker(host, ports, results)
	}

	go func() {
		for port := 1; port <= MAX_PORTS; port++ {
			ports <- port
		}
		close(ports)
	}()

	for i := 0; i < MAX_PORTS; i++ {
		if port := <-results; port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	return openPorts
}

func worker(host string, ports <-chan int, results chan<- int) {
	for port := range ports {
		address := net.JoinHostPort(host, strconv.Itoa(port))
		conn, err := net.DialTimeout("tcp", address, TIMEOUT)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- port
	}
}
