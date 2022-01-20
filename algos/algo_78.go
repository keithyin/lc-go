package algos

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	tracer := make([]int, 0)
	subsetsCore(nums, 0, &tracer, &result)
	return result
}

func subsetsCore(nums []int, pos int, tracer *[]int, result *[][]int) {
	if pos == len(nums) {
		tmp := make([]int, len(*tracer))
		copy(tmp, *tracer)
		*result = append(*result, tmp)
		return
	}

	subsetsCore(nums, pos+1, tracer, result)
	*tracer = append(*tracer, nums[pos])
	subsetsCore(nums, pos+1, tracer, result)
	*tracer = (*tracer)[:len(*tracer)-1]
	return

}
