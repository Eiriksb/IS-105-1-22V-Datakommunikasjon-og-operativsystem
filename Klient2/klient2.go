package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "8080"
	SERVER_TYPE = "tcp"
)

func caesar(r rune, shift int) rune {
	// Shift character by specified number of places.
	// ... If beyond range, shift backward or forward.
	s := int(r) + shift
	if s >= 248 { //Siste tegn i ascii vi bruker er 248 (å)
		return rune(s - 217) //Siste tegn - første tegn i ascii
	} else if s < 32 { //Første tegn i ascii vi bruker er 32 (space)
		return rune(s + 217)
	}
	return rune(s)
}

//fucntion to reacieve message from server
func recvLoop(r io.Reader) {
	var inbuf [1024]byte
	for {
		n, err := r.Read(inbuf[:])
		if err != nil {
			fmt.Println("Failed to receive msg from the server!")
			break
		}
		fmt.Print(string(inbuf[:n]))
	}
}

func main() {

	// Create a connection to the server.
	conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error connecting to server:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	// Receive messages from server
	for {
		go recvLoop(conn)
	}

	//response to string
	response := bufio.NewReader(os.Stdin)

	fmt.Println(response)
	fmt.Println()
	fmt.Println("----------------TILBAK FRA SERVER----------------")
	fmt.Println()
	fmt.Print("Server response: ")
	fmt.Println(response)
	// Close connection.
	conn.Close()

}
