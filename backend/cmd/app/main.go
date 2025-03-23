package main

import (
	"log"
	"net/http"

	"github.com/Cirillo-f/tcpScanner/configs"
)

func main() {
	config := configs.LoadConfig()
	rout := http.NewServeMux()

	server := http.Server{
		Addr:    config.PORT,
		Handler: rout,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("[ERROR]:", err)
	}

}
