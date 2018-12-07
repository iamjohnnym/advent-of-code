package main

import "testing"

func TestSum(t *testing.T) {
	total := sum([]float64{+1, +1, +1})
	if total != 3 {
		t.Errorf("Sum was incorrect, got: %f, want: %d.", total, 3)
	}
}
func TestSumTwo(t *testing.T) {
	total := sum([]float64{+1, +1, -2})
	if total != 0 {
		t.Errorf("Sum was incorrect, got: %f, want: %d.", total, 0)
	}
}

func TestSumThree(t *testing.T) {
	total := sum([]float64{-1, -2, -3})
	if total != -6 {
		t.Errorf("Sum was incorrect, got: %f, want: %d.", total, -6)
	}
}
