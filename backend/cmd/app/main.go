package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Cirillo-f/tcpScanner/configs"
	"github.com/Cirillo-f/tcpScanner/internal/handlers"
)

func main() {
	config := configs.LoadConfig()
	rout := http.NewServeMux()

	server := http.Server{
		Addr:    config.PORT,
		Handler: rout,
	}
	log.Println("[SUCCES]:Server has benn starting since ", time.Now())
	handlers.NewScanHandler(rout)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("[ERROR]:", err)
	}

}
