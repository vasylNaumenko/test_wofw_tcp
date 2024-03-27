package tcp_client_test

import (
	"net"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"sampisal/internal/repository/tcp_client"
)

func TestTCPClient_Read(t *testing.T) {
	conn1, conn2 := net.Pipe()
	defer conn1.Close()
	defer conn2.Close()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		_, err := conn1.Write([]byte("hello\n"))
		if err != nil {
			t.Fatal(err)
		}
	}()

	tcpClient := tcp_client.NewTCPClient().SetConn(conn2)

	data, err := tcpClient.Read()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "hello", data)
	wg.Wait()
}

func TestTCPClient_Write(t *testing.T) {
	conn1, conn2 := net.Pipe()
	defer conn1.Close()
	defer conn2.Close()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		tcpClient := tcp_client.NewTCPClient().SetConn(conn1)
		defer tcpClient.Close()

		err := tcpClient.Write("hello")
		if err != nil {
			t.Fatal(err)
		}
	}()

	tcpClient := tcp_client.NewTCPClient().SetConn(conn2)
	data, err := tcpClient.Read()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "hello", data)
	wg.Wait()
}
