package main

import (
	"flag"
	"time"

	"sampisal/internal/client/connector"
)

func main() {
	address := flag.String("address", "my-server", "The address of the server")
	port := flag.Uint("port", 8080, "The port of the server")
	delay := flag.Duration("delay", 1*time.Second, "The delay between each connection attempt")
	flag.Parse()

	connector.ConnectAndCommunicate(*address, uint16(*port), *delay)
}
