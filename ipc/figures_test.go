package ipc

import "testing"

// A status request may ask for the tunnel figures; no other command takes
// the field.
func TestOnlyStatusAsksForFigures(t *testing.T) {
	b, err := encode(&Request{Command: CmdStatus, Figures: true}, MaxRequest)
	if err != nil {
		t.Fatal(err)
	}
	if got, err := DecodeRequest(b); err != nil || !got.Figures {
		t.Fatalf("status asking for figures: %+v, %v", got, err)
	}
	b, err = encode(&Request{Command: CmdSelf, Figures: true}, MaxRequest)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := DecodeRequest(b); err == nil {
		t.Error("a self request asking for figures was accepted")
	}
}
