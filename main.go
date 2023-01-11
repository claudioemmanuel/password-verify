package main

import "github.com/claudioemmanue/go-api/pkg/server"

func main() {

	// Create a new server
	server := server.NewServer("8080")
	server.Start()
}
