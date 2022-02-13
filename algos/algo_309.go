package algos

/**
最佳买卖股票，含冷冻期。这个和打家劫舍有点像，但是又有些不一样，打家劫舍就是劫还是不劫。可以先思考暴力解法，然后通过剪枝的方式进行优化
这个如何进行暴力解法，然后进行优化呢？
*/

type State struct {
	state int // 0: empty, 1: 有股票在手，2：:处于冷冻期
}

func maxProfit309(prices []int) int {
	state := State{state: 0}
	cache := make(map[[2]int]int)
	return maxProfit309Core2(prices, 0, &state, cache)
}

func maxProfit309Core(prices []int, cur_pos int, state *State) int {
	if cur_pos >= len(prices) {
		return 0
	}
	if state.state == 0 { // 没有股票在手
		state.state = 1
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

func maxProfit309Core2(prices []int, cur_pos int, state *State, cache map[[2]int]int) int {
	if cur_pos >= len(prices) {
		return 0
	}

	if _, ok := cache[[2]int{state.state, cur_pos}]; ok {
		return cache[[2]int{state.state, cur_pos}]
	}

	val := -1

	if state.state == 0 { // 没有股票在手
		state.state = 1
		current_buy := maxProfit309Core2(prices, cur_pos+1, state, cache) - prices[cur_pos]
		state.state = 0
		current_not_buy := maxProfit309Core2(prices, cur_pos+1, state, cache)
		val = SliceMaxInt([]int{current_buy, current_not_buy})
	} else if state.state == 1 { // 有股票在手
		state.state = 2
		current_sold := maxProfit309Core2(prices, cur_pos+1, state, cache) + prices[cur_pos]
		state.state = 1
		current_not_sold := maxProfit309Core2(prices, cur_pos+1, state, cache)
		val = SliceMaxInt([]int{current_sold, current_not_sold})
	} else { // 当前处于冷冻期
		state.state = 0
		val = maxProfit309Core2(prices, cur_pos+1, state, cache)
		state.state = 2
	}

	cache[[2]int{state.state, cur_pos}] = val
	return val

}
