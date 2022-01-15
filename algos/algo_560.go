package algos

func subarraySum(nums []int, k int) int {
	cumSum := make([]int, len(nums)+1)
	cumSum[0] = 0
	for i := 1; i < len(cumSum); i++ {
		cumSum[i] = nums[i-1] + cumSum[i-1]
	}

	buffer := make(map[int]int, 0)
	buffer[0] = 1
	cnt := 0
	for i := 1; i < len(cumSum); i++ {
		if v, ok := buffer[cumSum[i]-k]; ok {
			cnt += v
		}

		if _, ok := buffer[cumSum[i]]; !ok {
			buffer[cumSum[i]] = 0
		}

		buffer[cumSum[i]] += 1
	}

	return cnt

}
