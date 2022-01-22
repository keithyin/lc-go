package algos

func maxProduct(nums []int) int {
	positiveMaxDp := make([]int, len(nums))
	negativeMaxDp := make([]int, len(nums))

	positiveMaxDp[0] = 1
	negativeMaxDp[0] = 1

	for i := 1; i < len(nums)+1; i++ {
		if nums[i-1] > 0 {
			positiveMaxDp[i] = positiveMaxDp[i-1] * nums[i-1]
			negativeMaxDp[i] = negativeMaxDp[i-1] * nums[i-1]
		} else if nums[i-1] < 0 {

		}
	}
}
