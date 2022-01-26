package event

import (
	"testing"
)

// Test som tester PutAllB
func TestKyoHsB(t *testing.T) {
	wanted := "Sts Ky:true | Sts Rev:false | Sts Korn:false | Sts HS:true"
	got := BoatInfo
	if got != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q. Enter er det ingen i båten, for mange i båten eller Mennesket mangler", got, wanted)
	} else {
		t.Log("Kylling og Menneske er i båten")
	}
}
