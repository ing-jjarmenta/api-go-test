package main

import (
	"log"

	"github.com/ing-jjarmenta/api-go-test/cmd/api/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal("Server stopped with error: ", err)
	}
}
