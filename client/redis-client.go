package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	_, error := conn.Write([]byte("Hello Server!"))
	if error != nil {
		fmt.Println(err)
		return
	}
}
