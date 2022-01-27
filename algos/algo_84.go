package algos

import "fmt"

func LargestRectangleArea(heights []int) int {
	return largestRectangleArea(heights)
}

func largestRectangleArea(heights []int) int {
	nextRightSmaller := make([]int, len(heights))
	stack := make([]int, 0)
	for i := len(heights) - 1; i >= 0; i-- {
		curVal := heights[i]
		for len(stack) > 0 && (curVal <= heights[stack[len(stack)-1]]) {
			stack = stack[0 : len(stack)-1]
		}
		if len(stack) == 0 {
			nextRightSmaller[i] = -1
		} else {
			nextRightSmaller[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = make([]int, 0)
	nextLeftSmaller := make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		curVal := heights[i]
		for len(stack) > 0 && (curVal <= heights[stack[len(stack)-1]]) {
			stack = stack[0 : len(stack)-1]
		}
		if len(stack) == 0 {
			nextLeftSmaller[i] = -1
		} else {
			nextLeftSmaller[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	fmt.Println(heights)
	fmt.Println(nextLeftSmaller)
	fmt.Println(nextRightSmaller)

	maxArea := 0
	for i := 0; i < len(heights); i++ {
		lefti := nextLeftSmaller[i]
		righti := nextRightSmaller[i]

		if righti == -1 {
			righti = len(heights)
		}
		area := (righti - lefti - 1) * heights[i]
		if maxArea < area {
			maxArea = area
		}
	}

	return maxArea

}
