package algos

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	results := make([][]int, 0)

	if len(intervals) > 0 {
		results = append(results, make([]int, 2))
		copy(results[0], intervals[0])
	}

	idx_j := 0
	for i := 1; i < len(intervals); i++ {
		if results[idx_j][1] >= intervals[i][0] {
			if results[idx_j][1] < intervals[i][1] {
				results[idx_j][1] = intervals[i][1]
			}
			if results[idx_j][0] > intervals[i][0] {
				results[idx_j][0] = intervals[i][0]
			}
		} else {
			results = append(results, make([]int, 2))
			copy(results[idx_j+1], intervals[i])
			idx_j++
		}
	}
	return results
}
