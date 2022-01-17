package algos

func hammingDistance(x int, y int) int {
	count := 0
	for i := 0; i < 32; i++ {
		if (x & 1) == (y & 1) {
			count += 1
		}
		x = x >> 1
		y = y >> 1
	}

	return count
}
