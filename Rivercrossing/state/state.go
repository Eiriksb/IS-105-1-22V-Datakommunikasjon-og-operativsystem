package state

import (
	"strconv"
)

//Disse skal kunne endres med Event

//Om de er på båten (Boat)
var KyB = false
var RevB = false
var KornB = false
var HsB = false

//Om de er på land V (West)
var KyW = true
var RevW = true
var KornW = true
var HsW = true
var BoatW = true

//om der er på land Ø (East)
var KyE = false
var RevE = false
var KornE = false
var HsE = false
var BoatE = false

//State på hvem som er i båt

func StateBoat() (FinalState string) {

	FinalState = "Sts Ky B:" + strconv.FormatBool(KyB) + " | " + "Sts Rev B:" + strconv.FormatBool(RevB) + " | " + "Sts Korn B:" + strconv.FormatBool(KornB) + " | " + "Sts HS B:" + strconv.FormatBool(HsB)

	return
}

//State på hvem som er på Land Vest

func StateLandV() (FinalState string) {

	FinalState = "Sts Ky V:" + strconv.FormatBool(KyW) + " | " + "Sts Rev V:" + strconv.FormatBool(RevW) + " | " + "Sts Korn V:" + strconv.FormatBool(KornW) + " | " + "Sts HS V:" + strconv.FormatBool(HsW) + " | " + "Sts Boat V: " + strconv.FormatBool(BoatW)

	return
}

//State på hvem som er på Land Øst

func StateLandE() (FinalState string) {

	FinalState = "Sts Ky E:" + strconv.FormatBool(KyE) + " | " + "Sts Rev E:" + strconv.FormatBool(RevE) + " | " + "Sts Korn E:" + strconv.FormatBool(KornE) + " | " + "Sts HS E:" + strconv.FormatBool(HsE) + " | " + "Sts Boat E: " + strconv.FormatBool(BoatE)

	return
}

//Legger sammen alle states i en string

func FinalState() (FinalState string) {
	FinalState = StateBoat() + StateLandE() + StateLandV()
	return
}

//For å så vise den med ViewState

func ViewState() string {
	return FinalState()
}
