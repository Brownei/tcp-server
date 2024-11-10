package main

import (
	"log"
	"net"
)

func main() {
	server := NewServer()
	go server.Run()

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Cannot connect to tcp: %v", err)
	}

	defer listener.Close()
	log.Printf("Connected to TCP server: %v", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Cannot be accepting the listener: %v", err)
			continue
		}

		go server.NewServerClient(conn)
	}
}
