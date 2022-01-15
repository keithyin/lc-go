package algos

func majorityElement(nums []int) int {
	majority := nums[0]
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if cnt == 0 {
			cnt += 1
			majority = nums[i]
			continue
		}

		if majority == nums[i] {
			cnt += 1
		} else {
			cnt -= 1
		}

	}
	return majority
}
