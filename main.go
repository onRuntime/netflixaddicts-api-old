package main

import "github.com/netflixaddicts/Go-API/server"

func main() {
	s := server.New()
	s.Initialize()
	s.Run(":8080")
}
