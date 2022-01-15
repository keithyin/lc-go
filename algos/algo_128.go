package algos

func longestConsecutive(nums []int) int {
	numsMap := make(map[int]bool, 0)
	for _, v := range nums {
		numsMap[v] = true
	}

	maxLen := 0

	for _, v := range nums {
		if _, ok := numsMap[v-1]; ok {
			continue
		}
		oriV := v
		for {
			if _, ok := numsMap[v+1]; ok {
				v += 1
			} else {
				break
			}
		}

		curLen := v - oriV + 1
		if curLen > maxLen {
			maxLen = curLen
		}

	}
	return maxLen
}
