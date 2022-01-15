package algos

func rob198(nums []int) int {
	return 0
}

func rob198Core(nums []int, i int, buffer map[int]int) int {
	if i >= len(nums) {
		return 0
	}

	if v, ok := buffer[i]; ok {
		return v
	}

	maxVal := SliceMaxInt([]int{nums[i] + rob198Core(nums, i+2, buffer), rob198Core(nums, i+1, buffer)})
	buffer[i] = maxVal
	return maxVal

}
