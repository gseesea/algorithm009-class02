package Week_04


// 最大利润
func maxProfit(prices []int) int {
	sum := 0
	for i := 0; i < len(prices) -1; i++ {
		// 有收益相加
		if prices[i+1] > prices[i] {
			sum += prices[i+1] - prices[i]
		}
	}
	return sum
}
