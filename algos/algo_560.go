package algos

func subarraySum(nums []int, k int) int {

	begin := 0
	end := 0

	cumulatedVal := nums[0]
	cnt := 0
	for begin < len(nums) && end < len(nums) {

		if begin > end {
			cumulatedVal += nums[end]
			end += 1
			continue
		}

		if cumulatedVal == k {
			cnt += 1
			cumulatedVal -= nums[begin]
			begin += 1
		} else if cumulatedVal > k {
			cumulatedVal -= nums[begin]
			begin += 1
		} else {
			cumulatedVal += nums[end]
			end += 1
		}
	}
	return cnt
}
