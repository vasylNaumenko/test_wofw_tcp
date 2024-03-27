package tcp_client

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"sampisal/internal/repository"
)

var _ repository.TCPClient = &TCPClient{}

type TCPClient struct {
	conn net.Conn
}

func NewTCPClient() *TCPClient {
	return &TCPClient{}
}

func (c *TCPClient) SetConn(conn net.Conn) *TCPClient {
	c.conn = conn
	return c
}

func (c *TCPClient) Dial(address string, port uint16) error {
	var err error
	c.conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", address, port))
	return err
}

func (c *TCPClient) Read() (string, error) {
	reader := bufio.NewReader(c.conn)
	data, err := reader.ReadString('\n')
	return strings.TrimSpace(data), err
}

func (c *TCPClient) Write(data string) error {
	_, err := c.conn.Write([]byte(data + "\n"))
	return err
}

func (c *TCPClient) Close() error {
	return c.conn.Close()
}

func (c *TCPClient) Listen(port uint, handlers ...func(client repository.TCPClient)) error {
	if len(handlers) == 0 {
		return fmt.Errorf("no handlers provided")
	}

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", uint16(port)))
	if err != nil {
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		client := &TCPClient{conn: conn}
		for _, handler := range handlers {
			go handler(client)
		}
	}

	return nil
}
