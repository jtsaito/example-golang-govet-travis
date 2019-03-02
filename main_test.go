package main

import "testing"

func TestSum(t *testing.T) {
	d := Sum(-1, 1)
	if d != 0 {
		t.Errorf("Sum -1 + 1 incorrect, got: %d, want: %d.", d, 0)
	}
}
