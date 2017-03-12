package leetcode

import "testing"

// Test461HammingDistance tests hammingDistance()
func Test461HammingDistance(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	distance := hammingDistance(123456, 1)

	if distance != 7 {
		t.Logf("Expected %v, but got %v", 7, distance)
		t.Fail()
	}
}
