package main

import (
	"fmt"
	"net"
)

func main() {
	// after connecting to the robot, found the ip address of the robot using the command:
	// echo "Current IP: $(ipconfig getifaddr en0)"
	conn, err := net.Dial("tcp", "192.168.4.2:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()
}