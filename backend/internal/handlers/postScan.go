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

	router.HandleFunc("/scan", handler.scanner())
}

func (handler *ScanHandler) scanner() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")              
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS") 
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")  

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		body, err := utils.HandleBody(&w, r)
		if err != nil {
			res.Json(w, err.Error(), http.StatusPaymentRequired)
			return
		}


		openPorts := scanner.TCPScanner(body.Host)


		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"host":       body.Host,
			"open_ports": openPorts,
		})
	}
}
