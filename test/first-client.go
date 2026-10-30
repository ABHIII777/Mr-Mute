package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("This is first client")

	fmt.Println("Trying to establish connection with the server")

	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("Some error occured:", err)
		return
	}

	fmt.Println("Connected to the server and sending the test message")

	defer conn.Close()

	var clientInput string

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the message: ")
		clientInput, _ = reader.ReadString('\n')
		clientInput = strings.TrimSpace(clientInput)

		if strings.Contains(clientInput, "exit") {
			return
		}
	
		data := []byte(clientInput)
		_, err = conn.Write(data)
	
		if err != nil {
			fmt.Println("Some error occured: ", err)
			return
		}
	}

	// data := []byte("Hello, This is a test message from first client")
	// _, err = conn.Write(data)

	// if err != nil {
	// 	fmt.Println("Some error occured:", err)
	// 	return
	// }
}
