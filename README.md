# TCP Chat Server in Go

A simple TCP chat server implemented in Go. This server allows multiple clients to connect and exchange messages in real-time.

---

## Features

- **Real-time messaging**: Clients can send and receive messages instantly.
- **Client management**: Automatically handles client connections and disconnections.
- **Concurrency**: Uses goroutines and channels for efficient handling of multiple clients.

---

## How It Works

- The server listens on a specified port for incoming client connections.
- Each client connection is handled in a separate goroutine.
- Messages from clients are broadcast to all other connected clients.
- The server uses a `Message` struct to manage client connections and messages.

---

## Installation

1. Clone the repository:
```bash
   git clone https://github.com/your-username/tcp-chat-server.git
   cd tcp-chat-server
   ```

2. Run the server:
```bash
   go run main.go
   ```

---
## Usage

1. Start the server:
```bash
go run main.go
```

2. Connect to the server using a TCP client (e.g., `telnet` or `nc`):
```
telnet localhost 6969
```

3. Start sending messages. All connected clients will receive the messages in real-time.

---

## Example

### Server Output

```bash 
# Server Output 
2023/10/10 12:00:00 Server is running on Port : 6969 
2023/10/10 12:00:05 client address 127.0.0.1:12345
```
### Client 1
```bash
# Client 1 
$ telnet localhost 6969 Trying 127.0.0.1... 
Connected to localhost. 
Hello, World!
```

### Client 2
```bash
# Client 2 
$ telnet localhost 6969 Trying 127.0.0.1... 
Connected to localhost. 
Hello, World!
```

---
## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

---

_This README was generated with the assistance of AI tools._
