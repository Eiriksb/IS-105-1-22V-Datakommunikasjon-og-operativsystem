package state

import (
	"strconv"
)

/*var ky bool
var rev bool
var korn bool
var hs bool
var boat bool
var water bool*/

func State() (FinalState string) {
	//Hvem som er på båten
	ky := true
	rev := true
	korn := true
	hs := true
	boat := true
	water := true
	//Returnerer status om den er true eller false og gjør den om til en string
	/*stateKy = "| Status Kylling: " + strconv.FormatBool(ky) + " |"
	stateRev = "Status Rev: " + strconv.FormatBool(rev) + " |"
	stateKorn = "Status Korn: " + strconv.FormatBool(korn) + " |"
	stateHS = "Status HS: " + strconv.FormatBool(hs) + " |"
	stateB = "Status Båt: " + strconv.FormatBool(boat) + " |"
	stateW = "Status Water: " + strconv.FormatBool(water) + " |"*/

	FinalState = "| Status Kylling: " + strconv.FormatBool(ky) + " |" + "Status Rev: " + strconv.FormatBool(rev) + " |" + "Status Korn: " + strconv.FormatBool(korn) + " |" + "Status HS: " + strconv.FormatBool(hs) + " |" + "Status Båt: " + strconv.FormatBool(boat) + " |" + "Status Water: " + strconv.FormatBool(water) + " |"

	return
}

func ViewState() string {
	return State()
}
