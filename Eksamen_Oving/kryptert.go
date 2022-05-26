package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var clientValue string

func caesar(r rune, shift int) rune {
	// Shift character by specified number of places.
	// ... If beyond range, shift backward or forward.
	s := int(r) + shift
	if s > 'å' {
		return rune(s - 29)
	} else if s < 'a' {
		return rune(s + 29)
	}
	return rune(s)
}

func stringToBin(s string) string {
	// Convert string to binary.
	var result string
	for _, r := range s {
		result += fmt.Sprintf("%08b", r)
	}
	return result
}

func main() {
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

	resultB := stringToBin(resultC)
	fmt.Println("Kryptert : " + resultC)
	fmt.Println("Binær : " + resultB)
}
