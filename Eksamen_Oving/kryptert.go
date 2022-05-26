package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var clientValue string

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "8080"
	SERVER_TYPE = "tcp"
)

func caesar(r rune, shift int) rune {
	// Shift character by specified number of places.
	// ... If beyond range, shift backward or forward.
	s := int(r) + shift
	if s > 'z' {
		return rune(s - 26)
	} else if s < 'a' {
		return rune(s + 26)
	}
	return rune(s)
}

func main() {

	// Connect to server.
	conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error connecting to server:", err.Error())
		os.Exit(1)

	}
	// Get client value.
	fmt.Print("Client til Server value: ")
	// Get input from user. (Simulated)
	scanner := bufio.NewScanner(os.Stdin)
	// Use `for scanner.Scan()` to keep reading
	scanner.Scan()
	clientValue = scanner.Text()
	fmt.Println("Tatt opp:", clientValue)
	// Test the caesar method in a func argument to strings.Map.

	resultC := strings.Map(func(r rune) rune {
		// Shift each character by 4 places.
		return caesar(r, 4)
	}, clientValue)

	fmt.Println("Kryptert og sendt : " + resultC)

	// Send client value to server.
	_, err = conn.Write([]byte(resultC))
	if err != nil {
		fmt.Println("Error sending to server:", err.Error())
		os.Exit(1)
	}
	// Read response from server.
	buffer := make([]byte, 1024)
	length, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from server:", err.Error())
		os.Exit(1)
	}
	// Print response from server
	response := []byte(buffer[:length])
	fmt.Println()
	fmt.Println("----------------TILBAK FRA SERVER----------------")
	fmt.Println()
	fmt.Print("Server response: ")
	fmt.Println(response)
	// Close connection.
	conn.Close()

	//Print response from server to a string
	resultD := strings.Map(func(r rune) rune {
		// Shift each character by 4 places.
		return caesar(r, -4)
	}, string(response))

	fmt.Println("Dekryptert og mottatt: " + resultD)
}
