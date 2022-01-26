package state

import (
	"testing"
)

func TestViewState(t *testing.T) {
	wanted := "| Status Kylling: true |Status Rev: true |Status Korn: true |Status HS: true |Status Båt: true |Status Water: true |"
	state := ViewState()
	if state != wanted {
		t.Errorf("Feil, fikk %q, ", state)
		t.Errorf("    ønsket %q", wanted)
	}
}
