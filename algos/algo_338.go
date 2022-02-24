package algos

import "math"

func countBits(n int) []int {
	bits := make([]int, n+1)
	for i := 0; i <= n; i++ {
		bits[i] = bitOneNum(i)
	}
	return bits
}

func bitOneNum(value int) int {
	if value == 0 {
		return 0
	}
	validBitsNum := int(math.Ceil(math.Log2(float64(value))) + 1)
	result := 0
	for i := 0; i < validBitsNum; i++ {
		result += ((value >> i) & 1)

	}
	return result
}
