package algos

func findUnsortedSubarray(nums []int) int {

	begin := 0
	end := len(nums) - 1
	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < max {
			end = i
		} else {
			max = nums[i]
		}
	}

	min := 9999

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > min {
			begin = i
		} else {
			min = nums[i]
		}
	}
	if end == begin {
		return 0
	} else {
		return end - begin + 1
	}
}
