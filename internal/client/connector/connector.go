package connector

import (
	"fmt"
	"strings"
	"time"

	"sampisal/internal/common/pow"
	"sampisal/internal/repository/tcp_client"
)

func ConnectAndCommunicate(address string, port uint16, delay time.Duration) {
	for {
		<-time.After(delay)
		err := communicateWithServer(address, port)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func communicateWithServer(address string, port uint16) error {
	client := tcp_client.NewTCPClient()
	err := client.Dial(address, port)
	if err != nil {
		return err
	}
	defer client.Close()

	challenge, err := client.Read()
	if err != nil {
		return fmt.Errorf("error reading challenge: %w", err)
	}
	challenge = strings.TrimSpace(challenge)
	fmt.Println("Received challenge:", challenge)

	response := pow.SolveChallenge(challenge)
	fmt.Println("Sending response:", response)

	err = client.Write(response)
	if err != nil {
		return fmt.Errorf("error sending response: %w", err)
	}

	quote, err := client.Read()
	if err != nil {
		return fmt.Errorf("error reading quote: %w", err)
	}

	fmt.Println("Received quote:", quote)
	return nil
}
