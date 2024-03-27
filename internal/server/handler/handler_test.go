package handler_test

import (
	"bufio"
	"net"
	"strings"
	"sync"
	"testing"

	"sampisal/internal/common/pow"
	"sampisal/internal/repository/tcp_client"
	"sampisal/internal/server/handler"
)

func TestHandleConnection(t *testing.T) {
	// Define the test cases
	testCases := []struct {
		name          string
		mockResponse  string
		expectedQuote bool
	}{
		{
			name:          "[success] Valid response to the challenge",
			mockResponse:  pow.SolveChallenge(pow.GenerateChallenge()),
			expectedQuote: true,
		},
		{
			name:          "[fail] Invalid response to the challenge",
			mockResponse:  "foobar",
			expectedQuote: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a pair of connected network connections
			conn1, conn2 := net.Pipe()
			defer conn1.Close()
			defer conn2.Close()

			var wg sync.WaitGroup
			wg.Add(1)

			// Write a response to the connection
			go func() {
				defer wg.Done()
				challenge, err := bufio.NewReader(conn1).ReadString('\n')
				if err != nil {
					t.Fatal(err)
				}
				challenge = strings.TrimSpace(challenge)

				if tc.expectedQuote {
					tc.mockResponse = pow.SolveChallenge(challenge)
				}

				writer := bufio.NewWriter(conn1)
				_, err = writer.WriteString(tc.mockResponse + "\n")
				if err != nil {
					t.Fatal(err)
				}
				writer.Flush()

				// Read the quote from the connection
				response, err := bufio.NewReader(conn1).ReadString('\n')
				if err != nil && tc.expectedQuote {
					t.Fatal(err)
				}

				if tc.expectedQuote {
					t.Log("Received quote:", response)
				}
			}()

			// Call the function with the mock connection
			handler.HandleConnection(tcp_client.NewTCPClient().SetConn(conn2))

			// Wait for the goroutine to finish
			wg.Wait()
		})
	}
}
