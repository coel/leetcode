package leetcode

func hammingDistance(x int, y int) int {
	xor := x ^ y
	return countBits(xor)
}

func countBits(i int) int {
	nbits := 0
	for i > 0 {
		nbits += i & 0x01
		i = i >> 1
	}
	return nbits
}

func countBits2(i int) int {
	i = (i & 0x55555555) + ((i >> 1) & 0x55555555)
	i = (i & 0x33333333) + ((i >> 2) & 0x33333333)
	i = (i & 0x0f0f0f0f) + ((i >> 4) & 0x0f0f0f0f)
	i = (i & 0x00ff00ff) + ((i >> 8) & 0x00ff00ff)
	i = (i & 0x0000ffff) + ((i >> 16) & 0x0000ffff)
	return i
}
