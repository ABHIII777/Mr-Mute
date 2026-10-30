package main

import (
	"fmt"
	"net"
)


func main() {
	fmt.Println("This is server !")

	listner, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("Some error occured:", err)
		return
	}

	defer listner.Close()
	
	fmt.Println("Server is listening on port 8080 !")

	for {
		conn, err := listner.Accept()

		if err != nil {
			fmt.Println("Some error occured:", err)
			return
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([] byte, 1024)
	for {
		n, err := conn.Read(buffer)

		if err != nil {
			fmt.Println("Some error occured:", err)
			return
		}

		fmt.Printf("Received: %s\n:", string(buffer[:n]))
	}
}
