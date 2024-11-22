package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server : ", err)
		os.Exit(1)
	}

	defer listener.Close()

	fmt.Println("Server is running on port 8080...")

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting connection : ", err)
			continue
		}

		go handleConnection(conn)

	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading from client", err)
		return
	}

	fmt.Printf("Received message : %s ", message)

	conn.Write([]byte("Message received \n"))
}
