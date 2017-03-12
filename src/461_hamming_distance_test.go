package leetcode

import "testing"

// TestHammingDistance tests hammingDistance()
func TestHammingDistance(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	hammingDistance(123456, 1)
}
