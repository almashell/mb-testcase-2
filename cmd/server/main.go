package main

import (
	"log"
	"mb_api/internal/app/server"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	server.Start()
}