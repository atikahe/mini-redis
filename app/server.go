package main

import (
	"fmt"
	"net"
	"os"
)

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)

	for {
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error reading from client")
			os.Exit(1)
		}

		conn.Write([]byte("+PONG\r\n"))
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

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("error accepting connection", err)
			continue
		}
		defer conn.Close()

		go handleRequest(conn)
	}

	// Create connection channel
	// events := make(chan net.Conn)

	// Receive connection
	// go func() {
	// 	for {
	// 		conn, err := l.Accept()
	// 		if err != nil {
	// 			fmt.Println("error accepting connection", err)
	// 			continue
	// 		}

	// 		// Register connection to events
	// 		events <- conn
	// 	}
	// }()

	// // Create new thread for new connection
	// for {

	// }

	// conn, err := l.Accept()
	// if err != nil {
	// 	fmt.Println("Error accepting connection: ", err.Error())
	// 	os.Exit(1)
	// }
	// defer conn.Close()

	// // Read multiple request
	// buf := make([]byte, 1024)
	// for {
	// 	_, err = conn.Read(buf)
	// 	if err != nil {
	// 		fmt.Println("error reading from client")
	// 		os.Exit(1)
	// 	}

	// 	conn.Write([]byte("+PONG\r\n"))
	// }
}
