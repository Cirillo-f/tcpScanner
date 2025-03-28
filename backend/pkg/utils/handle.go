package utils

import (
	"net/http"

	"github.com/Cirillo-f/tcpScanner/internal/models"
	"github.com/Cirillo-f/tcpScanner/pkg/res"
)

func HandleBody(w *http.ResponseWriter, r *http.Request) (*models.URL, error) {
	body, err := Decode(r.Body)
	if err != nil {
		res.Json(*w, err.Error(), http.StatusPaymentRequired)
		return nil, err
	}

	err = isValid(body.Host)
	if err != nil {
		res.Json(*w, err.Error(), http.StatusPaymentRequired)
		return nil, err
	}

	return body, nil
}
