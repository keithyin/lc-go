package algos

import "container/list"

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	var2Id := make(map[string]int)
	count := 0
	for _, equation := range equations {
		if _, ok := var2Id[equation[0]]; !ok {
			var2Id[equation[0]] = count
			count++
		}
		if _, ok := var2Id[equation[1]]; !ok {
			var2Id[equation[1]] = count
			count++
		}
	}

	edges := make([][][2]float64, count)
	for i := 0; i < count; i++ {
		edges = make([][][2]float64, 0)
	}
	for i, equation := range equations {
		first := var2Id[equation[0]]
		second := var2Id[equation[1]]
		edges[first] = append(edges[first], [2]float64{float64(second), values[i]})
		edges[second] = append(edges[second], [2]float64{float64(first), 1.0 / values[i]})
	}
	results := make([]float64, len(queries))
	for i, query := range queries {
		begin := var2Id[query[0]]
		end := var2Id[query[1]]
		if begin == end {
			results[i] = 1.0
		}

		visited := make([]bool, count)

		queue := list.New()
		queue.PushBack(begin)
		value := 1.0
		for queue.Len() > 0 {
			first := queue.Front().Value.(int)
			visited[first] = true

			for _, edge := range edges[first] {
				if visited[first] {
					continue
				}

				queue.PushBack(int(edge[0]))

			}

		}

	}

	return results

}
