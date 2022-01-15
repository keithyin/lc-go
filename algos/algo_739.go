package algos

import "container/list"

func dailyTemperatures(temperatures []int) []int {

	nextBiggerValueIdx := make([]int, len(temperatures))

	stack := list.New()

	for i := len(temperatures) - 1; i >= 0; i-- {
		for stack.Len() > 0 {
			if temperatures[stack.Back().Value.(int)] <= temperatures[i] {
				stack.Remove(stack.Back())
			} else {
				break
			}
		}

		if stack.Len() == 0 {
			nextBiggerValueIdx[i] = 0
		} else {
			nextBiggerValueIdx[i] = stack.Back().Value.(int) - i
		}

		stack.PushBack(i)

	}
	return nextBiggerValueIdx
}
