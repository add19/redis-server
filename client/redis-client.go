package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"resp"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	args := os.Args
	numArgs := len(args)
	var j resp.StringArray
	for i := 1; i < numArgs; i++ {
		j = append(j, args[i])
	}
	_, error := conn.Write([]byte(resp.Serialization(j)))

	connbuf := bufio.NewReader(conn)

	for {
		str, err := connbuf.ReadString('\n')
		if err != nil {
			break
		}

		if len(str) > 0 {
			fmt.Println(str)
		}
	}

	if error != nil {
		fmt.Println(err)
		return
	}
}
