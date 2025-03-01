package main

import (
	"fmt"
	"net"
	"time"
)

func testUDPPort(port int) {
	fmt.Printf("\n--- Testing UDP connection to port %d ---\n", port)
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("192.168.4.1:%d", port))
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Socket created, sending message")
	
	// Send a simple message
	message := []byte("HELLO\n")
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}
	fmt.Println("Data sent successfully")
	
	// Set a read deadline and try to receive response
	conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	buffer := make([]byte, 1024)
	
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err)
	} else {
		fmt.Printf("Received response: %s\n", buffer[:n])
	}
}

func testTCPPort(port int) {
	fmt.Printf("\n--- Testing TCP connection to port %d ---\n", port)
	addr := fmt.Sprintf("192.168.4.1:%d", port)
	
	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Successfully connected!")
	
	// Send a simple message
	_, err = conn.Write([]byte("HELLO\n"))
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}
	fmt.Println("Data sent successfully")
	
	// Try to read response
	buffer := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err)
	} else {
		fmt.Printf("Received response: %s\n", buffer[:n])
	}
}

func main() {
	// Test each port with both TCP and UDP
	ports := []int{80, 81, 100}
	
	for _, port := range ports {
		testTCPPort(port)
		testUDPPort(port)
	}
}