package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Cirillo-f/tcpScanner/pkg/res"
	"github.com/Cirillo-f/tcpScanner/pkg/scanner"
	"github.com/Cirillo-f/tcpScanner/pkg/utils"
)

type ScanHandler struct{}

func NewScanHandler(router *http.ServeMux) {
	handler := ScanHandler{}

	// Регистрируем обработчик для /scan
	router.HandleFunc("/scan", handler.scanner())
}

func (handler *ScanHandler) scanner() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Устанавливаем CORS заголовки на каждый запрос
		w.Header().Set("Access-Control-Allow-Origin", "*")              // Разрешаем доступ с любых доменов
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS") // Разрешаем методы POST и OPTIONS
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")  // Разрешаем Content-Type заголовки

		// Если это запрос OPTIONS, сразу завершаем его без дальнейшей обработки
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Быстрая проверка тела запроса для POST
		body, err := utils.HandleBody(&w, r)
		if err != nil {
			res.Json(w, err.Error(), http.StatusPaymentRequired)
			return
		}

		// Запускаем сканирование
		openPorts := scanner.TCPScanner(body.Host)

		// Отправляем ответ
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"host":       body.Host,
			"open_ports": openPorts,
		})
	}
}
