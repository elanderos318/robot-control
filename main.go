package main

import (
	"fmt"
	"net"
)

func main() {
	// after connecting to the robot, found the ip address of the robot using the command:
	// echo "Current IP: $(ipconfig getifaddr en0)"
	const robotAddress = "192.168.4.2:80"
	conn, err := net.Dial("udp", robotAddress)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	
	fmt.Println("Successfully connected to the robot!")

	defer conn.Close()
}