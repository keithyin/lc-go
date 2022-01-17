package algos

func findTargetSumWays(nums []int, target int) int {
	count := 0
	findTargetSumWaysCore(nums, 0, target, &count)
	return count
}

func findTargetSumWaysCore(nums []int, pos, target int, count *int) {
	if pos == len(nums) {
		if target == 0 {
			*count = *count + 1
		}
		return
	}

	findTargetSumWaysCore(nums, pos+1, target+nums[pos], count)
	findTargetSumWaysCore(nums, pos+1, target-nums[pos], count)

}
