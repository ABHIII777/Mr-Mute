package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	"mrmute/cmd/clientId"
	"mrmute/cmd/roomkey"
)

type Client struct {
	ID       string
	username string
	role     string
	Conn     net.Conn
}

type Room struct {
	ID           string
	Host         *Client
	Participants map[string]*Client
	JoinedAt     time.Time
	IsActive     bool
}

type Server struct {
	Rooms   map[string]*Room
	Clients map[string]*Client
	mu      sync.RWMutex
}

func main() {
	listner, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("Some error occured:", err)
		return
	}

	defer listner.Close()

	server := &Server{
		Rooms:   make(map[string]*Room),
		Clients: make(map[string]*Client),
	}

	fmt.Println("Server is online on port 8080 !")

	for {
		conn, err := listner.Accept()

		if err != nil {
			fmt.Println("Some error occured:", err)
			return
		}

		go server.handleClient(conn)
	}
}

func (server *Server) handleClient(conn net.Conn) {

	clientId, err := clientid.ClientID()

	if err != nil {
		fmt.Printf("Error occured: %s", err)
		return
	}

	client := &Client{
		ID:   clientId,
		Conn: conn,
	}

	server.mu.Lock()
	server.Clients[clientId] = client
	server.mu.Unlock()

	defer server.removeClient(clientId)
	defer conn.Close()

	fmt.Printf("Client connected: %s", clientId)

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)

		if err != nil {
			fmt.Println("Some error occured:", err)
			return
		}

		fmt.Printf("Received: %s\n:", string(buffer[:n]))

		clientResponse := string(buffer[:n])

		switch clientResponse {
		case "exit":
			_, err = conn.Write([]byte("Exiting..."))

			if err != nil {
				fmt.Println("Some error occured...")
				return
			}
			return

		case "key":
			key, err := roomkey.GenerateRoomKey()

			if err != nil {
				fmt.Printf("Some error occured while generating the room key")
				return
			}

			responseKey := []byte(key)
			_, err = conn.Write(responseKey)

			if err != nil {
				fmt.Printf("Error occured: %s", err)
				return
			}

			fmt.Printf("Responded with key: %s\n", key)
			continue

		default:
			_, err = conn.Write([]byte(clientResponse))

			if err != nil {
				fmt.Printf("Error occured: %s", err)
				return
			}

			continue
		}
	}
}

func (server *Server) removeClient(id string) {
	server.mu.Lock()
	defer server.mu.Unlock()

	if _, ok := server.Clients[id]; !ok {
		return
	}
	delete(server.Clients, id)

	for roomID, room := range server.Rooms {
		delete(room.Participants, id)

		if room.Host != nil && room.Host.ID == id {
			room.Host = nil
		}

		if len(room.Participants) == 0 && room.Host == nil {
			delete(server.Rooms, roomID)
		}
	}

	fmt.Printf("Client disconnected: %s\n", id)
}
