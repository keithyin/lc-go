package algos

func maxProduct(nums []int) int {
	maxValues := make([]int, len(nums))
	minValues := make([]int, len(nums))
	copy(maxValues, nums)
	copy(minValues, nums)
	for i := 1; i < len(nums); i++ {
		maxValues[i] = SliceMaxInt([]int{nums[i], nums[i] * maxValues[i-1], nums[i] * minValues[i-1]})
		minValues[i] = SliceMinInt([]int{nums[i], nums[i] * maxValues[i-1], nums[i] * minValues[i-1]})
	}
	return SliceMaxInt(maxValues)
}
