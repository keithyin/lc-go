package algos

import "fmt"

func FindKthLargest(nums []int, k int) int {
	Shuffle(nums)
	fmt.Println(nums)
	begin := 0
	end := len(nums) - 1
	idx := partitionCore(nums, begin, end)
	for idx != (k - 1) {
		fmt.Println(idx, nums)
		if idx > (k - 1) {
			end := idx - 1
			idx = partitionCore(nums, begin, end)
		} else {
			begin := idx + 1
			idx = partitionCore(nums, begin, end)
		}
	}

	fmt.Println(nums)

	return nums[idx]

}

func partitionCore(nums []int, begin, end int) int {

	if begin > end || end >= len(nums) {
		return -1
	}

	anchor := nums[begin]
	for begin < end {
		for anchor >= nums[end] && end > begin {
			end--
		}
		if end > begin {
			nums[begin] = nums[end]
			begin++
		} else {
			break
		}

		for anchor < nums[begin] && end > begin {
			begin++
		}
		if end > begin {
			nums[end] = nums[begin]
			end--
		} else {
			break
		}
	}

	nums[begin] = anchor

	return begin

}
