package repository

//go:generate mockgen -source repository.go -destination ./repository_mock.go -package repository

type (
	// TCPClient is an interface that defines the methods for TCP communication.
	TCPClient interface {
		// Dial establishes a connection to the server at the specified address and port.
		Dial(address string, port uint16) error

		// Listen starts a TCP server that listens on the specified address and port.
		// It uses the provided handlers to handle incoming connections.
		Listen(port uint, handlers ...func(client TCPClient)) error

		// Read reads data from the connection.
		// It returns the read data as a string and any error encountered.
		Read() (string, error)

		// Write sends the specified data over the connection.
		Write(data string) error

		// Close closes the connection.
		Close() error
	}
)
