package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Cirillo-f/tcpScanner/configs"
	"github.com/Cirillo-f/tcpScanner/internal/handlers"
	"github.com/Cirillo-f/tcpScanner/pkg/middleware"
)

func main() {
	config := configs.LoadConfig()
	rout := http.NewServeMux()
	handlers.NewScanHandler(rout)

	loggedRout := middleware.LoggerMiddleware(rout)

	server := http.Server{
		Addr:    config.PORT,
		Handler: loggedRout,
	}

	log.Println("[SUCCES]:Server has been starting since", time.Now())

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("[ERROR]:", err)
	}
}
