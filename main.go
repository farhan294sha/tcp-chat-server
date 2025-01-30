package main

import (
	"log"
	"net"
)

type MessageType int

const (
	ClientConnected MessageType = iota
	NewMessage
	DeleteClient
)

type Message struct {
	Type   MessageType
	Client net.Conn
	Text   string
}

const Port = "6969"

func server(message chan Message) {
	conns := make(map[string]net.Conn)

	for {
		msg := <-message

		switch msg.Type {
		case ClientConnected:
			conns[msg.Client.RemoteAddr().String()] = msg.Client
		case DeleteClient:
			delete(conns, msg.Client.RemoteAddr().String())
			msg.Client.Close()
		case NewMessage:
			for address, conn := range conns {
				if msg.Client == conn {
					continue
				}
				_, err := conn.Write([]byte(msg.Text))
				if err != nil {
					delete(conns, address)
					conn.Close()
				}
			}
		}
	}

}

func handleConnection(conn net.Conn, messages chan Message) {

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			messages <- Message{
				Type:   DeleteClient,
				Client: conn,
			}
			return
		}

		messages <- Message{
			Type:   NewMessage,
			Client: conn,
			Text:   string(buffer[:n]),
		}
	}

}

func main() {
	ln, err := net.Listen("tcp", ":"+Port)
	if err != nil {
		log.Fatalf("Failed to create Server %v", err)
	}
	log.Printf("Server is running on Port : %v", Port)

	messages := make(chan Message)
	go server(messages)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Failed to create connection %v", conn.RemoteAddr().String())
		}
		log.Printf("client address %v", conn.LocalAddr().String())
		messages <- Message{
			Type:   ClientConnected,
			Client: conn,
		}
		go handleConnection(conn, messages)
	}
}
