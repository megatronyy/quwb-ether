package tests

import (
	"testing"
)

func TestAdd(t *testing.T) {
	ret := Add(2, 3)
	if ret != 5 {
		t.Error("Expected 5, got ", ret)
	}
}

func TestMinus(t *testing.T) {
	ret := Minus(2, 3)
	if ret != -1 {
		t.Error("Expected -1, got ", ret)
	}
}
