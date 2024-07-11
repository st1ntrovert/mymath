package mymath

import (
	"testing"
)

func TestAbs(t *testing.T) {
	result := Abs(-1)
	expected := 1.0
	if result != expected {
		t.Errorf("Abs(-1) = %v; want %v", result, expected)
	}
}
