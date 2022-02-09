package algos

/**
最佳买卖股票，含冷冻期。这个和打家劫舍有点像，但是又有些不一样，打家劫舍就是劫还是不劫。可以先思考暴力解法，然后通过剪枝的方式进行优化
这个如何进行暴力解法，然后进行优化呢？
*/

type State struct {
	state int // 0: empty, 1: 有股票在手，2：:处于冷冻期
	val   int
}

func maxProfit309(prices []int) int {
	state := State{state: 0, val: 0}
	return maxProfit309Core(prices, 0, &state)
}

func maxProfit309Core(prices []int, cur_pos int, state *State) int {
	if cur_pos >= len(prices) {
		return 0
	}
	if state.state == 0 { // 没有股票在手
		state.state = 1
		state.val = prices[cur_pos]
		current_buy := maxProfit309Core(prices, cur_pos+1, state) - prices[cur_pos]
		state.state = 0
		current_not_buy := maxProfit309Core(prices, cur_pos+1, state)
		return SliceMaxInt([]int{current_buy, current_not_buy})
	} else if state.state == 1 { // 有股票在手
		state.state = 2
		current_sold := maxProfit309Core(prices, cur_pos+1, state) + prices[cur_pos]
		state.state = 1
		current_not_sold := maxProfit309Core(prices, cur_pos+1, state)
		return SliceMaxInt([]int{current_sold, current_not_sold})
	} else { // 当前处于冷冻期
		state.state = 0
		val := maxProfit309Core(prices, cur_pos+1, state)
		state.state = 2
		return val
	}
}
