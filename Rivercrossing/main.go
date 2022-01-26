package main

import (
	"bufio"
	"fmt"
	"os"

	"main.go/event"
	"main.go/state"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		fmt.Print("\033[H\033[2J")
		//Info Command
		if "info" == input.Text() {
			fmt.Println("Båt:       " + state.StateBoat())
			fmt.Println("Land Vest: " + state.StateLandV())
			fmt.Println("Land Øst:  " + state.StateLandE())
		}
		//Boat Commands
		if "PutBoatE" == input.Text() {
			event.PutBoatE()
			fmt.Println("Seilte Boat over til land Øst")
		}

		if "PutBoatV" == input.Text() {
			event.PutBoatV()
			fmt.Println("Seilte Boat over til land Vest")
		}
		//Kylling Commands
		if "PutKyB" == input.Text() {
			event.PutKyB()
			fmt.Println("Kylling er i båten")
		}

		if "PutKyV" == input.Text() {
			event.PutKyV()
			fmt.Println("Kylling er på Land Vest")
		}

		if "PutKyE" == input.Text() {
			event.PutKyE()
			fmt.Println("Kylling er på Land Øst")
		}
		//Rev Commands
		if "PutRevB" == input.Text() {
			event.PutRevB()
			fmt.Println("Rev er i båten")
		}

		if "PutRevV" == input.Text() {
			event.PutRevV()
			fmt.Println("Rev er på Land Vest")
		}

		if "PutRevE" == input.Text() {
			event.PutRevE()
			fmt.Println("Rev er på Land Øst")
		}
		//Korn Commands
		if "PutKornB" == input.Text() {
			event.PutKornB()
			fmt.Println("Korn er i båten")
		}

		if "PutKornV" == input.Text() {
			event.PutKornB()
			fmt.Println("Korn er på Land Vest")
		}

		if "PutKornE" == input.Text() {
			event.PutKornB()
			fmt.Println("Korn er på Land Øst")
		}
		//Hs Commands
		if "PutHsB" == input.Text() {
			event.PutHsB()
			fmt.Println("Menneske er i båten")
		}

		if "PutHsV" == input.Text() {
			event.PutHsV()
			fmt.Println("Menneske er på Land Vest")
		}

		if "PutHsE" == input.Text() {
			event.PutHsE()
			fmt.Println("Menneske er på Land Øst")
		}

		//Før over alle til de forskjellige plassene
		if "PutAllB" == input.Text() {
			event.PutAllB()
			fmt.Println("Alle er på Båten")
		}

		if "PutAllV" == input.Text() {
			event.PutAllV()
			fmt.Println("Alle er på Land Vest")
		}

		if "PutAllE" == input.Text() {
			event.PutAllE()
			fmt.Println("Alle er på Land Øst")
		} else {
			fmt.Println("Hva vil du?")

		}
	}

}
