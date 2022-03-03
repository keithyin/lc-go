package algos

func removeDuplicates(nums []int) int {
	anchor := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[anchor] {
			continue
		}
		anchor += 1
		nums[anchor] = nums[i]
	}
	return anchor + 1
}
