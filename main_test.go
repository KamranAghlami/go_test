package main

import "testing"

func TestSum(t *testing.T) {
	if Sum(2, 3) != 5 {
		t.Error("Expected 2 + 3 to equal 5")
	}
}
