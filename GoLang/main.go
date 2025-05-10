package main

import (
	"GoLang/utils"
	"bufio" // Helps to read data from the connection with buffering (like readline in Node).
	"fmt"
	"net" //Go's built-in networking package to work with TCP/UDP.
	"strings"
)

func main () {
	// Start listening on a port 
	listener, err := net.Listen("tcp", ":8080");
	// : Opens a TCP server on port 8080 (on all interfaces)

	// Error handling if the port is in use or invalid.
	if err != nil {
		fmt.Println("Error starting tcp server", err);
		return;
	}

	// Ensures the listener is closed when the program ends.
	defer listener.Close();

	fmt.Println("TCP server is listening on port: 8080");

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
	reader := bufio.NewReader(conn)
	for {
	//	fmt.Print(reader)
		value, err := utils.ParseRESP(reader);
		if err != nil {
			fmt.Println("Failed to parse:", err)
			break
		}
	
		
		fmt.Printf("Parsed RESP Value: %#v\n", value)

		parsedArray, ok := value.([]interface{})

		if ok {
			command := strings.ToUpper(parsedArray[0].(string));
			handler, exists := utils.Commands[command];

			if !exists {
				fmt.Println("Command not found:", command);
				conn.Write([]byte("-ERR unknown command\r\n"));
				continue;
			}
			args := make([]string, len(parsedArray[1:]))
			for i, v := range parsedArray[1:] {
				args[i] = v.(string)
			}
			result := handler(args);
			if result != nil {
				serialzieResult := utils.SerializeRESP(result);
				if serializedBytes, ok := serialzieResult.([]byte); ok {
					conn.Write(serializedBytes);
				} else {
					fmt.Println("Failed to serialize result to bytes");
					conn.Write([]byte("-ERR internal server error\r\n"));
				}
			}

		// Handle the RESP command, like PING, SET, GET...
	}
}


	fmt.Println("Client disconnected:", conn.RemoteAddr())
	}


