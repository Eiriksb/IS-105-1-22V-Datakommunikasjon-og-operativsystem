package event

import "main.go/state"

var BoatInfo = state.StateBoat()
var VestInfo = state.StateLandV()
var EastInfo = state.StateLandE()

//Tar Kylling i båt
func PutKyB() {
	state.KyB = true
	state.KyE = false
	state.KyW = false
}

//Tar Rev i båt
func PutRevB() {
	state.RevB = true
	state.RevE = false
	state.RevW = false
}

//Tar Korn i båt
func PutKornB() {
	state.KornB = true
	state.KornE = false
	state.KornW = false
}

//Tar Hs i båt
func PutHsB() {
	state.HsB = true
	state.HsE = false
	state.HsW = false
}

//Tar Kylling i Land Vest
func PutKyV() {
	state.KyB = false
	state.KyE = false
	state.KyW = true
}

//Tar Rev i Land Vest
func PutRevV() {
	state.RevB = false
	state.RevE = false
	state.RevW = true
}

//Tar Korn i Land Vest
func PutKornV() {
	state.KornB = false
	state.KornE = false
	state.KornW = true
}

//Tar Hs i Land Vest
func PutHsV() {
	state.HsB = false
	state.HsE = false
	state.HsW = true
}

//Tar Hs i Land Vest
func PutBoatV() {
	state.BoatE = true
	state.BoatW = false
}

//Tar Kylling i Land Øst
func PutKyE() {
	state.KyB = false
	state.KyE = true
	state.KyW = false
}

//Tar Rev i Land Vest
func PutRevE() {
	state.RevB = false
	state.RevE = true
	state.RevW = false
}

//Tar Korn i Land Vest
func PutKornE() {
	state.KornB = false
	state.KornE = true
	state.KornW = false
}

//Tar Hs i Land Vest
func PutHsE() {
	state.HsB = false
	state.HsE = true
	state.HsW = false
}

//Tar Hs i Land Vest
func PutBoatE() {
	state.BoatE = true
	state.BoatW = false
}

//Tar alle i Båten
func PutAllB() {
	state.KyB = true
	state.KyE = false
	state.KyW = false
	state.RevB = true
	state.RevE = false
	state.RevW = false
	state.KornB = true
	state.KornE = false
	state.KornW = false
	state.HsB = true
	state.HsE = false
	state.HsW = false
}

//Tar alle i Land Vest + båten
func PutAllV() {
	state.KyB = false
	state.KyE = false
	state.KyW = true
	state.RevB = false
	state.RevE = false
	state.RevW = true
	state.KornB = false
	state.KornE = false
	state.KornW = true
	state.HsB = false
	state.HsE = false
	state.HsW = true
	state.BoatE = false
	state.BoatW = true
}

//Tar alle i Land Øst + båten
func PutAllE() {
	state.KyB = false
	state.KyE = true
	state.KyW = false
	state.RevB = false
	state.RevE = true
	state.RevW = false
	state.KornB = false
	state.KornE = true
	state.KornW = false
	state.HsB = false
	state.HsE = true
	state.HsW = false
	state.BoatE = true
	state.BoatW = false
}
