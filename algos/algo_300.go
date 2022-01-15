package algos

import "fmt"

func LengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		curVal := nums[i]
		for j := i - 1; j >= 0; j-- {
			if curVal >= nums[j] {
				tmp := 1 + dp[j]
				if tmp > dp[i] {
					dp[i] = tmp
				}

			}
		}
	}
	fmt.Println(dp)
	return SliceMaxInt(dp)
}
