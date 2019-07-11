package main

import (
	"github.com/rumsrami/givly-rpc-api/http"
)

func main() {
	server := http.NewRPCServer()
	server.Run(":8080")
}
