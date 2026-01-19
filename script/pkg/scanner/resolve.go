package scanner

import (
	"fmt"
	"net"
)

func ResolveHost(host string) (string, error) {
	if ip := net.ParseIP(host); ip != nil {
		return host, nil
	}

	ips, err := net.LookupIP(host)
	if err != nil {
		return "", fmt.Errorf("failed to resolve %s: %w", host, err)
	}

	if len(ips) == 0 {
		return "", fmt.Errorf("no IP addresses found for %s", host)
	}

	for _, ip := range ips {
		if ip.To4() != nil {
			return ip.String(), nil
		}
	}

	return ips[0].String(), nil
}
