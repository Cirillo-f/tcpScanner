package utils

import (
	"errors"
	"regexp"
)

func isValid(data string) error {
	re := `^([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}$`
	if matched, _ := regexp.MatchString(re, data); !matched {
		return errors.New("[ERROR]: Invalid domain format.")
	}
	return nil
}
