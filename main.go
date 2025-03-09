package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"
)

type RobotCommand struct {
	N  int    `json:"N"`  // Command number
	H  string `json:"H"`  // Command serial number
	D1 int    `json:"D1"` // Direction (3 = forward)
	D2 int    `json:"D2"` // Speed (0-255)
	T  int    `json:"T"`  // Time in milliseconds
}

func main() {
	// Robot address - using the IP and port from your Wireshark capture
	robotAddr := "192.168.4.1:100"
	
	// Create a TCP connection
	conn, err := net.Dial("tcp", robotAddr)
	println(conn.RemoteAddr())
	if err != nil {
		log.Fatalf("Failed to connect to robot: %v", err)
	}
	defer conn.Close()

	// // Send GET /status request
	// _, err = conn.Write([]byte("GET /status HTTP/1.1\r\nHost: 192.168.4.1\r\n\r\n"))
	// if err != nil {
	// 	log.Fatalf("Failed to send status request: %v", err)
	// }

	// // Read response
	// newBuffer := make([]byte, 4096)
	// conn.SetReadDeadline(time.Now().Add(time.Second * 5))
	// n, err := conn.Read(newBuffer)
	// if err != nil {
	// 	log.Printf("No response or error: %v", err)
	// } else {
	// 	fmt.Printf("Status response: %s\n", string(newBuffer[:n]))
	// }

	
	fmt.Println("Connected to robot successfully")
	
	// Create a command to move forward for 1 second
	// cmd := RobotCommand{
	// 	N:  2,         // Command 2 = Car control with time limit
	// 	H:  "Go-001",  // Command serial number
	// 	D1: 3,         // Direction 3 = Forward
	// 	D2: 150,       // Speed = 150
	// 	T:  10000,      // Time = 1000ms (1 second)
	// }

	cmd := RobotCommand{
		N:  2,         // Command 2 = Car control with time limit
		H:  "Go-001",  // Command serial number
		D1: 3,         // Direction 3 = Forward
		D2: 150,       // Speed = 150
		// T:  10000,      // Time = 1000ms (1 second)
	}
	
	// Convert command to JSON
	jsonData, err := json.Marshal(cmd)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	
	// Send the command
	fmt.Printf("Sending command: %s\n", string(jsonData))
	_, err = conn.Write(jsonData)
	if err != nil {
		log.Fatalf("Failed to send command: %v", err)
	}
	
	// Read response
	buffer := make([]byte, 4096)
	conn.SetReadDeadline(time.Now().Add(time.Second * 5))
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("No response or error: %v", err)
	} else {
		fmt.Printf("Response from robot: %s\n", string(buffer[:n]))
	}
	
	fmt.Println("Command sent successfully")
}