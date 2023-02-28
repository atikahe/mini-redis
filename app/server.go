package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/atikahe/mini-redis/pkg/resp"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// buf := make([]byte, 1024)

	for {
		// _, err := conn.Read(buf)
		// if err != nil {
		// 	fmt.Println("error decoding from client")
		// 	os.Exit(1)
		// }

		input, err := resp.Decode(bufio.NewReader(conn))
		if err != nil {
			fmt.Println("error decoding from client")
			os.Exit(1)
		}

		fmt.Println("INPUT", input)

		conn.Write([]byte("+PONG\r\n"))

		// command := input.Array()[0].String()
		// args := input.Array()[1:]

		// switch command {
		// case "ping":
		// 	fmt.Println("+PONG\r")
		// 	conn.Write([]byte("+PONG\r\n"))
		// case "echo":
		// 	// Handle echo
		// 	fmt.Println("args", args)
		// default:
		// 	conn.Write([]byte("-ERR unknown command '" + command + "'\r\n"))
		// }

	}

}

func main() {
	fmt.Println("Logs from your program will appear here!")

	// Bind TCP server to port 6379
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		fmt.Println("CONNECTION ACCEPTED")
		if err != nil {
			fmt.Println("error accepting connection", err)
			continue
		}

		go handleConnection(conn)
	}
}
