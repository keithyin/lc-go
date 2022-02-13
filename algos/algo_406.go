package algos

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		} else {
			return people[i][0] > people[j][0]
		}
	})

	res := make([][]int, 0)
	for _, pair := range people {
		if pair[1] >= len(res) {
			res = append(res, pair)
		} else {
			tmp := make([][]int, 0)
			tmp = append(tmp, pair)
			tmp = append(tmp, res[pair[1]:]...)
			res = append(res[:pair[1]], tmp...)
		}
	}
	return res

}
