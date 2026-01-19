package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Cirillo-f/tcpScanner/pkg/scanner"
	"github.com/Cirillo-f/tcpScanner/pkg/utils"
)

const helpText = `tcpScanner - simple TCP port scanner

Usage:
  tcpScanner -d <domain> [-a | -p]

Flags:
  -d string    Domain or IP address to scan (required)
  -a           Scan all 65535 ports
  -p           Scan only popular ports
  -h           Show this help message

Examples:
  tcpScanner -d example.com -p
  tcpScanner -d 192.168.1.1 -a
`

var spinnerChars = []string{"|", "/", "-", "\\"}

func main() {
	var (
		domain   = flag.String("d", "", "Domain or IP address to scan")
		allPorts = flag.Bool("a", false, "Scan all 65535 ports")
		popPorts = flag.Bool("p", false, "Scan only popular ports")
		help     = flag.Bool("h", false, "Show help message")
	)

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, helpText)
	}

	flag.Parse()

	if *help {
		fmt.Print(helpText)
		os.Exit(0)
	}

	if *domain == "" {
		fmt.Fprintf(os.Stderr, "Error: domain required, use -d flag\n")
		fmt.Fprintf(os.Stderr, "Use -h for help\n")
		os.Exit(1)
	}

	if !*allPorts && !*popPorts {
		fmt.Fprintf(os.Stderr, "Error: must specify either -a (all ports) or -p (popular ports)\n")
		fmt.Fprintf(os.Stderr, "Use -h for help\n")
		os.Exit(1)
	}

	if *allPorts && *popPorts {
		fmt.Fprintf(os.Stderr, "Error: cannot use -a and -p simultaneously\n")
		os.Exit(1)
	}

	if err := utils.ValidateDomain(*domain); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Fprintf(os.Stderr, "\n\nInterrupted, stopping scan...\n")
		cancel()
	}()

	sc := scanner.New()

	var openPorts []int
	var totalPorts int

	if *allPorts {
		totalPorts = 65535
		fmt.Printf("Starting scan for %s (all 65535 ports)\n", *domain)
	} else {
		totalPorts = len(scanner.PopularPorts)
		fmt.Printf("Starting scan for %s (%d popular ports)\n", *domain, totalPorts)
	}

	done := make(chan bool)
	var mu sync.Mutex
	scanned := 0
	progressChan := make(chan int, 1000)

	go func() {
		spinnerIndex := 0
		ticker := time.NewTicker(50 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				mu.Lock()
				current := scanned
				mu.Unlock()
				spinnerIndex = (spinnerIndex + 1) % len(spinnerChars)
				percentage := float64(current) / float64(totalPorts) * 100
				fmt.Printf("\rScanning... %s [%d/%d] %.1f%%", spinnerChars[spinnerIndex], current, totalPorts, percentage)
			}
		}
	}()

	go func() {
		for range progressChan {
			mu.Lock()
			scanned++
			mu.Unlock()
		}
	}()

	var err error
	if *allPorts {
		openPorts, err = sc.ScanAll(ctx, *domain, progressChan)
	} else {
		openPorts, err = sc.ScanPopular(ctx, *domain, progressChan)
	}

	close(progressChan)
	time.Sleep(100 * time.Millisecond)
	close(done)
	time.Sleep(100 * time.Millisecond)

	fmt.Print("\r")
	fmt.Printf("Done! Scanned %d port(s)\n\n", totalPorts)

	if err != nil {
		if ctx.Err() == context.Canceled {
			fmt.Println("Scan cancelled")
			os.Exit(130)
		}
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if len(openPorts) == 0 {
		fmt.Println("No open ports found")
		os.Exit(0)
	}

	fmt.Printf("\nFound %d open port(s):\n", len(openPorts))
	fmt.Println("PORT     STATE  SERVICE")
	fmt.Println("----------------------")
	for _, port := range openPorts {
		service := scanner.GetServiceName(port)
		if service != "unknown" {
			fmt.Printf("%-8d open   %s\n", port, service)
		} else {
			fmt.Printf("%-8d open   (unknown)\n", port)
		}
	}
}
