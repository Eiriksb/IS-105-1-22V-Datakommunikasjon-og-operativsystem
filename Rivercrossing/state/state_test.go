package state

import (
	"strings"
	"testing"
)

//Tester om alle er her
func TestBoat(t *testing.T) {

	isTrue := "true"

	allOnBoard := "Sts Ky:true | Sts Rev:true | Sts Korn:true | Sts HS:true | Sts Wa: true"

	state := StateBoat()
	if state != allOnBoard {
		t.Log("Noen er på båten")
	} else if strings.Contains(state, isTrue) {
		t.Log("Alle er på båten")
	} else {
		t.Errorf("Ingen er på båten %q, ", state)
	}
}
func TestLandV(t *testing.T) {

	isTrue := "true"

	allOnBoard := "Sts Ky:true | Sts Rev:true | Sts Korn:true | Sts HS:true | Sts Wa: true"

	state := StateLandV()
	if state != allOnBoard {
		t.Log("Noen er på land Vest")
	} else if strings.Contains(state, isTrue) {
		t.Log("Alle er på land Vest")
	} else {
		t.Errorf("Ingen er på land Vest %q, ", state)
	}
}

func TestLandE(t *testing.T) {

	isTrue := "true"

	allOnBoard := "Sts Ky:true | Sts Rev:true | Sts Korn:true | Sts HS:true | Sts Wa: true"

	state := StateLandE()
	if state != allOnBoard {
		t.Log("Noen er på land Øst")
	} else if strings.Contains(state, isTrue) {
		t.Log("Alle er på land Øst")
	} else {
		t.Errorf("Ingen er på land Øst %q, ", state)
	}
}
