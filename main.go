package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.4.2:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()
}