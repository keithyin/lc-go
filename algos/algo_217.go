package algos

func containsDuplicate(nums []int) bool {
	existMap := make(map[int]struct{}, 0)
	for _, v := range nums {
		if _, ok := existMap[v]; ok {
			return true
		}
		existMap[v] = struct{}{}
	}
	return false
}
