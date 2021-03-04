package main

import (
	"github.com/hi-hi-ray/desafio-sw-go/api/handler"
	"log"
)

func main() {
	log.Println("Starting API server")
	handler.HandleRequests()
}
