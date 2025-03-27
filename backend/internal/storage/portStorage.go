package storage

import "sync"

var (
	PORTsave []string
	mu       sync.Mutex
)

// Очистка перед новым сканированием
func Clear() {
	mu.Lock()
	PORTsave = nil
	mu.Unlock()
}

// Сохранение портов без дубликатов
func SavePorts(portSet map[string]struct{}) {
	mu.Lock()
	PORTsave = make([]string, 0, len(portSet))
	for port := range portSet {
		PORTsave = append(PORTsave, port)
	}
	mu.Unlock()
}
