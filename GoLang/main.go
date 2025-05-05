package main

import (
	"GoLang/utils"
	"bufio" // Helps to read data from the connection with buffering (like readline in Node).
	"fmt"
	"net" //Go's built-in networking package to work with TCP/UDP.
)

func main () {
	// Start listening on a port 
	listener, err := net.Listen("tcp", ":8000");
	// : Opens a TCP server on port 8080 (on all interfaces)

	// Error handling if the port is in use or invalid.
	if err != nil {
		fmt.Println("Error starting tcp server", err);
		return;
	}

	// Ensures the listener is closed when the program ends.
	defer listener.Close();

	fmt.Println("TCP server is listening on port: 8000");

	for {
		//Accepts new clients one by one in a loop like socket in node.js
		// conn is a net.Conn object
		conn, err := listener.Accept();
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		// Handle the connection in a new goroutine
		// Spawns a goroutine to handle each client concurrently
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close();

	fmt.Println("Client Connected...", conn.RemoteAddr());
	
	//Creates a line-by-line reader.
	scanner := bufio.NewScanner(conn);

	for scanner.Scan() {

		text := scanner.Text();
		fmt.Println("Recieved", text);

		var args, err = utils.ParserRESP([]byte(text))
		fmt.Println(args);
		//Echo back the message Sends back a message to the client.
		//_,err := conn.Write([]byte("Echo:" + text + "\n"));
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