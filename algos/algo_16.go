package algos

func threeSumClosest(nums []int, target int) int {
	currMin := 0xfffffff
	currClosest := -1
	for i := 0; i < len(nums)-2; i++ {
		residualMap := make(map[int]struct{}, 0)
		for j := i + 1; j < len(nums); j++ {
			if len(residualMap) > 0 {
				for residual := range residualMap {
					if absInt(residual-nums[j]) < currMin {
						currMin = absInt(residual - nums[j])
						currClosest = target - (residual - nums[j])
					}
				}
			}
			residualMap[target-nums[i]-nums[j]] = struct{}{}
		}
	}
	return currClosest
}

func absInt(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}
