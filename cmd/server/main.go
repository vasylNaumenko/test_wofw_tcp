package main

import (
	"flag"

	"sampisal/internal/repository/tcp_client"
	"sampisal/internal/server/handler"
)

func main() {
	port := flag.Uint("port", 8080, "The port of the server")
	flag.Parse()

	tcpClient := tcp_client.NewTCPClient()
	err := tcpClient.Listen(*port, handler.HandleConnection)
	if err != nil {
		panic(err)
	}
}
