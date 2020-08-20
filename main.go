package main

import (
	"github.com/netflixaddicts/Go-API/server"
	"log"
)

func main() {
	log.Print("Starting program...")
	s := server.New()
	s.Initialize()
	s.Run(":8080")
}
