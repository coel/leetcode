package leetcode

func findComplement(num int) int {
	mask := getUsedBitsMask(num)
	return ^num & mask
}

func getUsedBitsMask(num int) int {
	mask := 0

	if num == 0 {
		return mask
	}

	for num > 1 {
		num = num >> 1
		mask = mask | 0x01
		mask = mask << 1
	}

	mask = mask | 0x01

	return mask
}
