package algos

func maxProfit(prices []int) int {
	minVal := 9999999999
	maxP := -99999
	for _, v := range prices {
		if v < minVal {
			minVal = v
		}
		if maxP < (v - minVal) {
			maxP = v - minVal
		}
	}
	return maxP
}
