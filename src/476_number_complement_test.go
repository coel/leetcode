package leetcode

import "testing"

// Test476FindComplement tests findComplement()
func Test476FindComplement(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	complement := findComplement(123456)

	if complement != 0x1DBF {
		t.Logf("Expected %v, but got %v", 0x1DBF, complement)
		t.Fail()
	}
}

// Test476MaskUsedBits tests getUsedBitsMask()
func Test476MaskUsedBits(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	mask := getUsedBitsMask(1234)

	if mask != 0x7FF {
		t.Logf("Expected %v, but got %v", 0x7FF, mask)
		t.Fail()
	}
}
