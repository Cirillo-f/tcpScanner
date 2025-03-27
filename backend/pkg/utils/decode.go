package utils

import (
	"encoding/json"
	"io"

	"github.com/Cirillo-f/tcpScanner/internal/models"
)

func Decode(body io.ReadCloser) (*models.URL, error) {
	var payload models.URL
	defer body.Close()

	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}
