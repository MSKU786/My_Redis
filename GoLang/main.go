package main

import (
	"bufio"
	"fmt"
	"net"
)

func main () {
	// Start listening on a port 
	listener, err := net.Listen("tcp", ":8000");
	if err != nil {
		fmt.Println("Error starting tcp server", err);
		return;
	}

	defer listener.Close();

	fmt.Println("TCP server is listening on port: 8000");

	for {
		//Aceept a connection
		conn, err := listener.Accept();
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close();

	fmt.Println("Client Connected...", conn.RemoteAddr());
	scanner := bufio.NewScanner(conn);

	for scanner.Scan() {
		text := scanner.Text();
		fmt.Println("Recieved", text);

		//Echo back the message
		_,err := conn.Write([]byte("Echo:" + text + "\n"));
		if (err != nil) {
			fmt.Println("Error writting to client", err);
			return;
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Connection error:", err)
	}
	fmt.Println("Client disconnected:", conn.RemoteAddr())

}