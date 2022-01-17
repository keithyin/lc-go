package algos

func topKFrequent(nums []int, k int) []int {

	counter := make(map[int]int, 0)
	for _, v := range nums {
		if _, ok := counter[v]; !ok {
			counter[v] = 0
		}
		counter[v] += 1
	}

	return []int{1}

}
