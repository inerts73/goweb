package mymath

import (
	"testing"
)

func TestAddInt(t *testing.T) {
	got := AddInt(2,3)
	if got != 5 {
		t.Errorf("WRONG")
	}
}

func TestMultiplyInt(t *testing.T) {
	got := MultiplyInt(2,3)
	if got != 6 {
		t.Errorf("WRONG")
	}
}