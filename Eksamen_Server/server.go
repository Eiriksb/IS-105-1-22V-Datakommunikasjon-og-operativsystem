package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "8080"
	SERVER_TYPE = "tcp"
)

func main() {
	fmt.Println("Server Running...")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()
	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for client...")
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("client connected")
		go processClient(connection)
	}
}

//Bytes to binary string
func toBinaryBytes(s string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(s); i++ {
		fmt.Fprintf(&buffer, "%08b", s[i])
	}
	return fmt.Sprintf("%s", buffer.Bytes())
}

func processClient(connection net.Conn) {
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	//serverResponse := string(buffer[:mLen])
	response := toBinaryBytes(string(buffer[:mLen]))

	fmt.Println("Received : ", string(buffer[:mLen]))
	fmt.Println("Sendes back in binary : ", string(response))
	_, err = connection.Write([]byte(response))
	fmt.Println("Sent in byte: ", []byte(response))
	connection.Close()
}
