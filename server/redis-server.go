package main

import (
	"fmt"
	"net"
	"resp"
)

type data_structure struct {
	dict map[string]string
}

var data data_structure

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
		var j resp.String = "PONG"
		conn.Write([]byte(resp.Serialization(j)))
	case "ECHO":
		var j resp.String = resp.String(command[1])
		conn.Write([]byte(resp.Serialization(j)))
	case "SET":
		key := command[1]
		val := command[2]
		data.dict[key] = val
		var j resp.String = "OK"
		conn.Write([]byte(resp.Serialization(j)))
	case "GET":
		key := command[1]
		v, ok := data.dict[key]
		if !ok {
			conn.Write([]byte("$-1\r\n"))
		} else {
			var j resp.String = resp.String(v)
			conn.Write([]byte(resp.Serialization(j)))
		}
	default:
		conn.Write([]byte(stringResp))
	}
}

func main() {

	l, err := net.Listen("tcp", "localhost:6379")
	if err != nil {
		return
	}

	data = data_structure{
		dict: make(map[string]string), // Initialize the map
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
