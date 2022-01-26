package state

import (
	"testing"
)

//Tester om alle er her
func TestBoat(t *testing.T) {

	wanted := "Sts Ky B:true | Sts Rev B:true | Sts Korn B:true | Sts HS B:true"

	state := StateBoat()
	if state == wanted {
		t.Log("Alle er på Båt")
	} else {
		t.Errorf("Ingen er på båten %q, ", state)
	}
}
func TestLandV(t *testing.T) {
	wanted := "Sts Ky V:true | Sts Rev V:true | Sts Korn V:true | Sts HS V:true | Sts Boat V: true"

	var state = StateLandV()
	if state == wanted {
		t.Log("Alle er på land Vest")
	} else {
		t.Errorf("Ingen er på land Vest %q, ", state)
	}
}

func TestLandE(t *testing.T) {

	wanted := "Sts Ky E:true | Sts Rev E:true | Sts Korn E:true | Sts HS E:true | Sts Boat E: true"

	state := StateLandE()
	if state == wanted {
		t.Log("Alle er på land Øst")
	} else {
		t.Errorf("Ingen er på land Øst %q, ", state)
	}
}
