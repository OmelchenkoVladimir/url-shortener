package encoding

import (
	"testing"
)

func TestWrongSymbols(t *testing.T) {
	_, err := Decode("hello")
	if err == nil {
		t.Errorf("String containts wrong symbols but function still decoded it")
	}
}

func TestEncode(t *testing.T) {
	d, _ := Decode("hi")
	e := Encode(d)
	if e != "hi" {
		t.Errorf("Decoded %s, want hi; btw number was %d", e, d)
	}
}
