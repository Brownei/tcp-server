# tcp-server

# TCP Server

This project implements a simple TCP server in Golang to handle client connections, receive data, and respond to requests. The server is designed for basic networking tasks such as communication with clients over TCP sockets, sending responses, and handling multiple clients concurrently.

## Features

- **TCP Socket Communication**: The server accepts connections from clients over TCP sockets, allowing bidirectional communication.
- **Multi-Client Support**: The server can handle multiple client connections.
- **Configurable**: Server parameters, such as IP address, port, and buffer size, can be configured.
- **Logging**: Logs incoming connections, data received, and responses sent for easy monitoring.

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/brownei/tcp-server.git
    cd tcp-server
    ```

2. Install any required dependencies:

    ```bash
    go mod tidy
    ```

### Usage

1. **Run the Server**:

    ```bash
    go run cmd/*.go
    ```

2. **Connect Clients**: Use any TCP client (e.g., `telnet`, or a custom client) to connect to the server. For example:

    ```bash
    telnet localhost 8000
    ```

### Example

Here's a basic example of client-server interaction:

1. Client connects to the server.
2. Server acknowledges the connection.
3. Client sends a message.
4. Server processes and responds to the message.

## Contributing

I will really appreciate any other upgrade in the project

---

Happy Coding!
