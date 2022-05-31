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
	if s >= 248 { //Siste tegn i ascii vi bruker er 248 (å)
		return rune(s - 217) //Siste tegn - første tegn i ascii
	} else if s < 32 { //Første tegn i ascii vi bruker er 32 (space)
		return rune(s + 217)
	}
	return rune(s)
}

//Convert binary into ascii text
func binaryToString(data []byte) string {
	var result string
	for _, value := range data {
		result += string(value)
	}
	return result
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
	mlen, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from server:", err.Error())
		os.Exit(1)
	}
	//serverResponse := string(buffer[:mLen])
	response := string(buffer[:mlen])
	fmt.Println()
	fmt.Println("----------------TILBAK FRA SERVER----------------")
	fmt.Println()
	fmt.Print("Server response: ")
	fmt.Println(response)
	// Close connection.
	conn.Close()

	// Decrypt response
	resultD := strings.Map(func(r rune) rune {
		// Shift each character by 4 places.
		return caesar(r, -4)
	}, string(response))

	fmt.Println("Dekryptert og mottatt: " + resultD)
}
