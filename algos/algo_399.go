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
		edges[i] = make([][2]float64, 0)
	}
	for i, equation := range equations {
		first := var2Id[equation[0]]
		second := var2Id[equation[1]]
		edges[first] = append(edges[first], [2]float64{float64(second), values[i]})
		edges[second] = append(edges[second], [2]float64{float64(first), 1.0 / values[i]})
	}
	results := make([]float64, len(queries))
	for i, query := range queries {
		if _, ok := var2Id[query[0]]; !ok {
			results[i] = -1.0
			continue
		}
		if _, ok := var2Id[query[1]]; !ok {
			results[i] = -1.0
			continue
		}
		begin := var2Id[query[0]]
		end := var2Id[query[1]]
		if begin == end {
			results[i] = 1.0
			continue
		}

		queue := list.New()
		queue.PushBack(begin)
		tmpValues := make([]float64, count)
		for i := 0; i < count; i++ {
			tmpValues[i] = -1.0
		}
		tmpValues[begin] = 1.0
		for queue.Len() > 0 {
			firstNode := queue.Front()
			first := firstNode.Value.(int)
			queue.Remove(firstNode)
			for _, edge := range edges[first] {
				curIdx := int(edge[0])
				if tmpValues[curIdx] != -1.0 {
					continue
				}
				tmpValues[curIdx] = tmpValues[first] * edge[1]
				if curIdx == end {
					break
				}
				queue.PushBack(curIdx)

			}
			if tmpValues[end] != -1.0 {
				break
			}
		}
		results[i] = tmpValues[end]

	}

	return results

}
