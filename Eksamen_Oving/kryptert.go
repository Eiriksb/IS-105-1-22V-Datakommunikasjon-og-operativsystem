package main

import (
	"fmt"
	"strings"
)

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

func main() {
	// Test the caesar method in a func argument to strings.Map.
	value := "w, x og y møtes i ålesund"
	result := strings.Map(func(r rune) rune {
		return caesar(r, 4)
	}, value)
	fmt.Println("Ikke Kryptert : " + value)
	fmt.Println("Kryptert : " + result)
}
