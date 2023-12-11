package main

import (
	"fmt"
	"net"
)

func handleConnectionRequest(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)

	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Received: %s\n", buf)
}

func main() {

	l, err := net.Listen("tcp", "localhost:6379")
	if err != nil {
		return
	}

	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnectionRequest(c)
	}
}
