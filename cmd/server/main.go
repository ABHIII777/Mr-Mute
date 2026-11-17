package main

import (
	"fmt"
	"log"
	"mrmute/cmd/roomkey"
	"net"
	"time"
)

type Client struct {
	ID       string
	username string
	role     string
	conn     net.Conn
}

type Room struct {
	ID           string
	Host         *Client
	Participants map[string]*Client
	JoinedAt     time.Time
	IsActive     bool
}

type Server struct {
	Rooms map[string]*Room
}

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

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)

		if err != nil {
			fmt.Println("Some error occured:", err)
			return
		}

		fmt.Printf("Received: %s\n:", string(buffer[:n]))

		if string(buffer[:n]) == "key" {
			key, err := roomkey.GenerateRoomKey()

			if err != nil {
				fmt.Println("Error occured while generating the room key !")
			}

			responseKey := []byte(key)
			_, err = conn.Write(responseKey)

			if err != nil {
				log.Fatal(err)
				return
			}
		}
	}
}
