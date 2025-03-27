package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/Cirillo-f/tcpScanner/internal/storage"
)

func TcpScanner(url string) {
	var wg sync.WaitGroup
	portSet := make(map[string]struct{}) // Используем map для исключения дубликатов
	var mu sync.Mutex

	storage.Clear() // Очистка перед новым сканированием

	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			address := net.JoinHostPort(url, fmt.Sprintf("%d", port))
			conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
			if err != nil {
				return
			}
			conn.Close()

			mu.Lock()
			portSet[fmt.Sprintf("%d", port)] = struct{}{}
			mu.Unlock()
		}(i)

		if i%100 == 0 {
			time.Sleep(50 * time.Millisecond)
		}
	}
	wg.Wait()

	// Конвертируем map в срез перед сохранением
	storage.SavePorts(portSet)
}
