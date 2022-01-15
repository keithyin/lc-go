package algos

import "container/list"

func maxSlidingWindow(nums []int, k int) []int {

	results := make([]int, 0)

	queue := list.New()

	begin := 0

	for i := 0; i < len(nums); i++ {
		lastEle := queue.Back()
		for lastEle != nil {
			if nums[lastEle.Value.(int)] <= nums[i] {
				queue.Remove(lastEle)
				lastEle = queue.Back()
			} else {
				break
			}
		}
		queue.PushBack(i)

		if (i - begin + 1) > k {
			begin += 1
		}

		firstEle := queue.Front()
		if firstEle != nil {
			if firstEle.Value.(int) < begin {
				queue.Remove(firstEle)
			}
		}

		if (i - begin + 1) == k {
			results = append(results, nums[i])
		}
	}
	return results
}
