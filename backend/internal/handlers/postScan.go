package handlers

import (
	"net/http"

	"github.com/Cirillo-f/tcpScanner/internal/storage"
	"github.com/Cirillo-f/tcpScanner/pkg/res"
	"github.com/Cirillo-f/tcpScanner/pkg/scanner"
	"github.com/Cirillo-f/tcpScanner/pkg/utils"
)

type ScanHandler struct{}

func NewScanHandler(router *http.ServeMux) {
	handler := ScanHandler{}

	router.HandleFunc("POST /scan", handler.scanner())
}

func (handler *ScanHandler) scanner() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := utils.HandleBody(&w, r)
		if err != nil {
			return
		}

		scanner.TcpScanner(body.URL)
		w.Header().Set("Content-Type", "application/json")
		res.Json(w, storage.PORTsave, http.StatusOK)
	}
}
