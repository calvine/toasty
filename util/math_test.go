package util

import "testing"

func TestMin(t *testing.T) {
	x := Min(4, 1)
	if x != 1 {
		t.Error("wrong answer for min")
	}
}

func TestMax(t *testing.T) {
	x := Max(4, 1)
	if x != 4 {
		t.Error("wrong answer for max")
	}
}
