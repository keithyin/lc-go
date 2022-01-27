package algos

func moveZeroes(nums []int) {

	a, b := 0, 0
	for ; b < len(nums); b++ {
		for a < len(nums) && nums[a] != 0 {
			a++
		}
		if nums[b] != 0 && a < b {
			tmp := nums[a]
			nums[a] = nums[b]
			nums[b] = tmp
		}
	}
}
