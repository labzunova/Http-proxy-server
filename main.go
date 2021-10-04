package main

import (
	"log"
	"net/http"

	"proxy/proxy.go/proxy"
)

func main() {
	handler := &proxy.HttpProxy{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
