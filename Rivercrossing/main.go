package main

import (
	"fmt"

	"main.go/state"
)

func main() {
	fmt.Println("Hva vil du?: putIn | putOut | crossRiver")
	fmt.Println("Liste over hvor ting er akkurat nå: ")
	fmt.Println("Båt:       " + state.StateBoat())
	fmt.Println("Land Vest: " + state.StateLandV())
	fmt.Println("Land Øst:  " + state.StateLandE())
}
