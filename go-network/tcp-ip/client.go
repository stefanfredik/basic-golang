package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		fmt.Println("Error connecting to server", err)
		os.Exit(1)
	}

	defer conn.Close()

	fmt.Print("Enter message to send : ")
	reader := bufio.NewReader(os.Stdin)
	message, _ := reader.ReadString('\n')

	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending data : ", err)
		return
	}

	response, err := bufio.NewReader(conn).ReadString('\n')

	if err != nil {
		fmt.Println("Error receiving response : ", err)
		return
	}

	fmt.Println("Server response : %s ", response)
}
