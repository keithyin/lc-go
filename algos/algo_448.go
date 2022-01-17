package algos

import "fmt"

func FindDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for (i+1) != nums[i] && nums[nums[i]-1] != nums[i] {
			val := nums[i]
			tmp := nums[val-1]
			nums[i] = tmp
			nums[val-1] = val
		}

	}

	fmt.Println(nums)

	results := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if (i + 1) != nums[i] {
			results = append(results, i+1)
		}
	}

	return results
}
