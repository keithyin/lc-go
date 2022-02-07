package algos

import "fmt"

func TopKFrequent(nums []int, k int) []int {
	return topKFrequent(nums, k)
}

func topKFrequent(nums []int, k int) []int {

	counter := make(map[int]int)
	for _, v := range nums {
		if _, ok := counter[v]; !ok {
			counter[v] = 0
		}
		counter[v] += 1
	}

	tuples := dict2tuple(counter)
	fmt.Println(tuples)
	begin := 0
	end := len(tuples) - 1

	for begin <= end {
		curPos := tuplePartition(tuples, begin, end)
		if curPos == (k - 1) {
			result := make([]int, k)
			for i := 0; i <= curPos; i++ {
				result[i] = tuples[i][0]
			}
			return result
		} else if curPos > (k - 1) {
			end = curPos - 1
		} else {
			begin = curPos + 1
		}

	}

	return make([]int, 0)

}

func dict2tuple(counter map[int]int) [][]int {
	results := make([][]int, len(counter))
	for i := 0; i < len(results); i++ {
		results[i] = make([]int, 2)
	}

	idx := 0

	for k, v := range counter {
		results[idx] = []int{k, v}
		idx++
	}
	return results
}

func tuplePartition(values [][]int, begin, end int) int {
	anchor := values[begin]
	for begin < end {
		for end > begin && values[end][1] < anchor[1] {
			end--
		}
		if end > begin {
			values[begin] = values[end]
			begin++
		}

		for end > begin && values[begin][1] >= anchor[1] {
			begin++
		}
		if end > begin {
			values[end] = values[begin]
			end--
		}

	}
	values[begin] = anchor
	return begin
}
