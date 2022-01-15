package algos

func sortColors(nums []int) {

	zeroPos := 0
	onePos := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			tmp := nums[zeroPos]
			nums[zeroPos] = 0
			nums[i] = tmp
			zeroPos += 1
		}
	}
	onePos = zeroPos
	for i := zeroPos; i < len(nums); i++ {
		if nums[i] == 1 {
			tmp := nums[onePos]
			nums[onePos] = 1
			nums[i] = tmp
			onePos += 1
		}
	}
}
