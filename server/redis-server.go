package main

import (
	"fmt"
	"net"
	"resp"
)

func handleConnectionRequest(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	stringResp := string(buf[:])

	if err != nil {
		fmt.Println(err)
		return
	}

	command := resp.Deserialization(stringResp)

	switch command[0] {
	case "PING":
		conn.Write([]byte("PONG\n"))
	case "ECHO":
		conn.Write([]byte(command[1] + "\n"))
	default:
		conn.Write([]byte(stringResp))
	}
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
