package utils

import (
	"errors"
	"net"
	"regexp"
)

func ValidateDomain(domain string) error {
	if domain == "" {
		return errors.New("domain cannot be empty")
	}

	if ip := net.ParseIP(domain); ip != nil {
		return nil
	}

	re := `^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`
	if matched, _ := regexp.MatchString(re, domain); !matched {
		return errors.New("invalid domain or IP address format")
	}

	return nil
}
