package state

import (
	"strconv"
)

//Disse skal kunne endres med Event

//Om de er på båten (Boat)
var kyB = true
var revB = false
var kornB = false
var hsB = false

//Om de er på land V (West)
var kyW = false
var revW = false
var kornW = true
var hsW = false
var boatW = true

//om der er på land Ø (East)
var kyE = false
var revE = true
var kornE = false
var hsE = true
var boatE = false

//State på hvem som er i båt

func StateBoat() (FinalState string) {

	FinalState = "Sts Ky:" + strconv.FormatBool(kyB) + " | " + "Sts Rev:" + strconv.FormatBool(revB) + " | " + "Sts Korn:" + strconv.FormatBool(kornB) + " | " + "Sts HS:" + strconv.FormatBool(hsB)

	return
}

//State på hvem som er på Land Vest

func StateLandV() (FinalState string) {

	FinalState = "Sts Ky:" + strconv.FormatBool(kyW) + " | " + "Sts Rev:" + strconv.FormatBool(revW) + " | " + "Sts Korn:" + strconv.FormatBool(kornW) + " | " + "Sts HS:" + strconv.FormatBool(hsW) + " | " + "Sts Boat: " + strconv.FormatBool(boatW)

	return
}

//State på hvem som er på Land Øst

func StateLandE() (FinalState string) {

	FinalState = "Sts Ky:" + strconv.FormatBool(kyE) + " | " + "Sts Rev:" + strconv.FormatBool(revE) + " | " + "Sts Korn:" + strconv.FormatBool(kornE) + " | " + "Sts HS:" + strconv.FormatBool(hsE) + " | " + "Sts Boat: " + strconv.FormatBool(boatE)

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
