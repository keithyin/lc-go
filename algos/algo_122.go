package algos

func maxProfit122(prices []int) int {
	return maxProfit122Core(prices, 0, 0)
}

func maxProfit122Core(prices []int, i int, curState int) int {
	// curState: 1 has a coupon, 0. don't have coupon
	if i >= len(prices) {
		return 0
	}

	if curState == 0 {
		curBuy := maxProfit122Core(prices, i+1, 1) - prices[i]
		curNotBuy := maxProfit122Core(prices, i+1, 0)
		return SliceMaxInt([]int{curBuy, curNotBuy})
	} else {
		curSell := maxProfit122Core(prices, i+1, 0) + prices[i]
		curNotSell := maxProfit122Core(prices, i+1, 1)
		return SliceMaxInt([]int{curSell, curNotSell})
	}
}
