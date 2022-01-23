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

func dict2tuple(counter map[int]int) [][]int {
	results := make([][]int, len(counter))
	for i := 0; i < len(results); i++ {
		results[i] = make([]int, 2)
	}

	idx := 0

	for k, v := range counter {
		results[idx] = []int{k, v}
	}
	return results
}

func kthLargest(values [][]int, begin, end int) {

	anchor := values[begin]
	for begin < end {
		for values[end][1] <             
	}

}
