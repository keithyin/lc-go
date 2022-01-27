package algos

func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != (i + 1) {
			idx := nums[i]
			curPosNewVal := nums[idx-1]
			if curPosNewVal == idx {
				return curPosNewVal
			}
			nums[idx-1] = idx
			nums[i] = curPosNewVal
		}
	}
	return -1
}
